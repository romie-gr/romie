//nolint:gosec,noctx
package emulatorgames

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"

	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/utils"
)

var (
	emulatorGamesValidConsoles = []string{"playstation"}
	collectedGames             []string
	collectedPages             []string
)

func Parse(console string) {
	defer utils.TimeTrack(time.Now(), "EmulatorGames Parser")

	if !utils.StringContains(console, emulatorGamesValidConsoles...) {
		log.Panicf("Platform %s is not yet supported", console)
	}

	uri := fmt.Sprintf("https://www.emulatorgames.net/roms/%s/", console)

	for _, paginationURL := range collectPaginationLinks(uri) {
		parseListPage(paginationURL)
	}

	for _, gameURL := range collectedGames {
		parseGame(gameURL, console)
		log.Printf("Game: %s", gameURL)
	}
}

func parseAndGetDocument(uri string) *goquery.Document {
	// Make HTTP request
	response, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Errorf("Error loading HTTP response body (%v)", err)
		return nil
	}

	return document
}

func parseListPage(uri string) {
	document := parseAndGetDocument(uri)
	document.Find("a.eg-box").Each(processGameLink)
}

func processGameLink(_ int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		collectedGames = append(collectedGames, href)
	}
}

func collectPaginationLinks(uri string) []string {
	// Add current page, so even if there are no pagination links
	// the current page will be returned
	_ = append(collectedPages, uri)

	document := parseAndGetDocument(uri)
	document.Find("a.page-link").Each(processPaginationLink)

	return collectedPages
}

// Finds and collects unique pagination links
func processPaginationLink(_ int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists && href != "#" {
		collectedPages = appendIfMissing(collectedPages, href)
	}
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}

	return append(slice, i)
}

// GetDownloadLink returns the direct download link for a given game URL
func GetDownloadLink(gameURL string) (downloadLink string, err error) {

	// Create a temp directory (why? Because it starts downloading automatically after 10 seconds)
	var dir string
	if dirPath, err := os.Getwd(); err != nil {
		log.Println(err)
	} else {
		dir, err = ioutil.TempDir(dirPath, "chromedp-example")
		if err != nil {
			panic(err)
		}
	}
	defer os.RemoveAll(dir) // remove the directory (including the half-finished downloaded file)

	// Create our custom context background
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// Set the timeout limit as par the context
	ctx, cancel = context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	// Open the browser using the previously created context
	_ = chromedp.Run(ctx)

	// Browser settings
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-ipc-flooding-protection", false),
		chromedp.Flag("disable-client-side-phishing-detection", false),
		chromedp.Flag("disable-background-timer-throttling", false),
		chromedp.WindowSize(1200, 800),
		chromedp.Flag("headless", false), // just for debugging this code
		chromedp.Flag("hide-scrollbars", false),
		// chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	)

	// Apply the browser settings into the browser
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	// Life is not perfect. This code is triggered when there's either a timeout or error occurs
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer func() {
		fmt.Println("close browser")
		cancel()
	}()

	// clicky-here-clicky-there-wait-a-bit-and-click-over-there
	var ok bool
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(gameURL),
		chromedp.Click("/html/body/div[3]/div[2]/div[3]/form[1]/button"),
		chromedp.ActionFunc(logAction("Save Game is clicked")),
		chromedp.Sleep(time.Millisecond*600),
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(dir), // download the file into a specfic dir
		chromedp.WaitVisible("/html/body/div[3]/div[2]/div[1]/p/span[2]/a"),
		chromedp.ActionFunc(logAction("Download link is available")),
		chromedp.AttributeValue("/html/body/div[3]/div[2]/div[1]/p/span[2]/a", "href", &downloadLink, &ok),
	)

	return strings.TrimSpace(downloadLink), err
}

func logAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Info(logStr)
		return nil
	}
}

func parseGame(gameURL string, console string) {
	document := parseAndGetDocument(gameURL)
	name, _ := document.Find("h1[itemprop='name']").Html()
	lang, _ := document.Find(".eg-meta").Html()
	downloadLink, err := GetDownloadLink(gameURL)
	if err != nil {
		downloadLink = "n/a"
	}

	rom := scraper.NewRom(name, console, lang, gameURL, downloadLink)

	log.Println(rom.Stringer())
}
