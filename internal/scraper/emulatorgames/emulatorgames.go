//nolint:gosec,noctx
package emulatorgames

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
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

func parseGame(gameURL string, console string) {
	document := parseAndGetDocument(gameURL)
	name, _ := document.Find("h1[itemprop='name']").Html()
	lang, _ := document.Find(".eg-meta").Html()

	rom := scraper.NewRom(name, console, lang, gameURL, "")

	log.Println(rom.Stringer())
}
