package emulatorgames

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/page"
	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/utils"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

func parsePageGames(uri string) {
	document := utils.ParseAndGetDocument(uri)
	document.Find("a.site-box").Each(processGameLink)
}

func processGameLink(_ int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		pageURL, _ := url.Parse(href)
		if err := utils.IsOnline(*pageURL); err != nil {
			log.Fatalf("page %s is not accessible: %v", pageURL, err)
		}
		log.Debugf("Page %s is accessible.", pageURL)
		collectedGames = append(collectedGames, href)
	}
}

func parseGame(gameURL string, console string) {
	document := utils.ParseAndGetDocument(gameURL)
	name, _ := document.Find("h1[itemprop='name']").Html()
	lang, _ := document.Find("div[class='site-post-meta mb-2']").Html()
	downloadLink, err := getDownloadLink(gameURL)

	if err != nil {
		downloadLink = "n/a"
	}

	rom := scraper.NewRom(name, console, lang, gameURL, downloadLink)
	Roms = append(Roms, *rom)

	log.Println(rom.Stringer())
}

// getDownloadLink returns the direct download link for a given game URL
// nolint:funlen
func getDownloadLink(gameURL string) (downloadLink string, err error) {
	// Create a temp directory (why? Because it starts downloading automatically after 10 seconds)
	var tempDownloadPath string

	if dirPath, err := os.Getwd(); err != nil {
		log.Println(err)
	} else {
		tempDownloadPath, err = ioutil.TempDir(dirPath, "chromedp-example")
		if err != nil {
			panic(err)
		}
	}

	// remove the directory (including the half-finished downloaded file)
	defer os.RemoveAll(tempDownloadPath)

	// Create a custom context background
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// Set the timeout limit as part the context
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
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		// chromedp.DisableGPU,
		chromedp.UserDataDir(tempDownloadPath),
	)

	// Apply browser settings to chromedp instance
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	// Life is not perfect. This code is triggered when there's either a timeout or error occurs
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	defer func() {
		log.Debug("close browser")
		cancel()
	}()

	// click-here-click-there-wait-a-bit-and-click-over-there
	var ok bool
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(gameURL),
		chromedp.Click("/html/body/div[3]/div[2]/div[3]/form[1]/button"),

		chromedp.ActionFunc(LogAction("Clicked on 'Save Game'")),
		chromedp.Sleep(time.Millisecond*600),

		// download the file into a specific tempDownloadPath
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(tempDownloadPath),
		chromedp.WaitVisible("/html/body/div[3]/div[2]/div[1]/p/span[2]/a"),

		chromedp.ActionFunc(LogAction("Found download link")),
		chromedp.AttributeValue("/html/body/div[3]/div[2]/div[1]/p/span[2]/a", "href", &downloadLink, &ok),
	)

	//nolint:wrapcheck
	return strings.TrimSpace(downloadLink), err
}

func LogAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Info(logStr)
		return nil
	}
}
