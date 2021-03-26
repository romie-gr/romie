package emulatorgames

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/romie-gr/romie/internal/utils"
	log "github.com/sirupsen/logrus"
)

func getPaginationLinks(uri string) []string {
	pageURL, _ := url.Parse(uri)
	if err := utils.IsOnline(*pageURL); err != nil {
		log.Fatalf("page %s is not accessible: %v", pageURL, err)
	}

	log.Debugf("Page %s is accessible.", pageURL)

	document := utils.ParseAndGetDocument(uri)
	log.Debugf("Searching for 'a.page-link' at %s", uri)
	document.Find("a.page-link").Each(processPaginationLink)

	for k, v := range collectedPages {
		if k == 0 {
			collectedPages[k] = uri
		} else {
			link := fmt.Sprintf("%v%s/", uri, v)
			pageURL, _ := url.Parse(link)
			if err := utils.IsOnline(*pageURL); err != nil {
				log.Fatalf("page %s is not accessible: %v", pageURL, err)
			}
			log.Debugf("Page %s is accessible.", pageURL)
			collectedPages[k] = link
		}
	}

	return collectedPages
}

// Finds and collects unique pagination links
func processPaginationLink(counter int, element *goquery.Selection) {
	lastPageNumber := 0
	href, exists := element.Attr("href")
	log.Debugf("a.page-link href [%d]: %v", counter+1, href)

	if exists && href != "#" {
		title := element.Text()
		if strings.Contains(title, "...") {
			sliceLink := strings.Split(href, "/")
			s := sliceLink[len(sliceLink)-2]
			// string to int
			lastPageNumber, _ = strconv.Atoi(s)
		} else {
			log.Debug("No ... found")
		}
	} else {
		log.Debugf("No valid href")
	}

	if lastPageNumber != 0 {
		log.Debugf("Valid link: %v", href)

		for i := 1; i <= lastPageNumber; i++ {
			collectedPages = append(collectedPages, fmt.Sprintf("%d", i))
		}
	}
}
