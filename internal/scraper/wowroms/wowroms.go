package wowroms

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/scraper/emulatorgames"
	"github.com/romie-gr/romie/internal/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

const wowroms = "https://wowroms.com"

var console string

func Parse(_console string) {
	defer utils.TimeTrack(time.Now(), "Wowroms Parser")

	// TODO MAKE CLASS?
	console = _console
	url := fmt.Sprintf("%s/en/roms/list/%s/", wowroms, console)

	document := emulatorgames.ParseAndGetDocument(url)

	pageNum := getNumberOfPages(document)

	// TODO GOROUTINIZE
	for i := 1; i <= pageNum; i++ {
		// Create url that points to this page number
		pageUrl := fmt.Sprintf("%s?page=%d", url, i)

		// Parse each page
		parsePage(pageUrl)
	}
}

func parsePage(url string) {
	var gameUrls []string

	document := emulatorgames.ParseAndGetDocument(url)

	document.Find("a.hoverBorder").Each(func(i int, s *goquery.Selection) {
		gameLink, _ := s.Attr("href")
		gameUrls = append(gameUrls, wowroms+gameLink)
	})

	// TODO GOROUTINIZE
	for _, gameUrl := range gameUrls {
		parseGame(gameUrl)
	}

}

func parseGame(url string) {
	//fmt.Println(url)
	document := emulatorgames.ParseAndGetDocument(url)

	downloadRef, _ := document.Find("a.btn.btn-prev.col-md-24.btnDwn").Attr("href")
	downloadLink := wowroms + downloadRef
	name := strings.TrimSpace(document.Find("div.col-md-24 b").Text())
	rom := scraper.NewRom(name, console, "", url, downloadLink)
	log.Println(rom.Stringer())

	// downloadRom(downloadLink)
}

func getDownloadLink(url string) string {
	splitUrl := strings.Split(url, "/")
	return fmt.Sprintf("%s%s/download-%s", splitUrl[0], console, splitUrl[1])
}

func downloadRom(url string) {
	done := make(chan bool)

	// create chrome instance
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if evt, ok := ev.(*page.EventDownloadProgress); ok {
			fmt.Printf("current download state is %s\n", evt.State.String())
			if evt.State == page.DownloadProgressStateCompleted {
				done <- true
			}
		}
	})

	chromedp.Run(ctx, chromedp.Tasks{
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath("."),
		chromedp.Navigate(url),
	})
	<-done
	fmt.Println("download finished")
}

func getNumberOfPages(document *goquery.Document) int {
	href := ""

	// Iterate through all the pagination links.
	document.Find("a.btn.alphabetP").Each(func(i int, s *goquery.Selection) {
		href, _ = s.Attr("href")
	})

	// The last pagination link contains the index of the last page (the number of pages)
	pageNum, err := strconv.Atoi(strings.Split(href, "=")[1])

	if err != nil {
		log.Errorf("Error converting number of pages to string: %v", err)
	}

	return pageNum
}
