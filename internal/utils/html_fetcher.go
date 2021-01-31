package utils

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//GetHTMLObj returns the HTML document for a user provided url
func GetHTMLObj(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery is unable to parse the returned html")
	}
	return doc, nil
}
