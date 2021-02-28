package utils

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetHTML returns the HTML document for a user provided url
func GetHTML(url string) (*goquery.Document, error) {
	/* #nosec G107: Potential HTTP request made with variable url */
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	defer res.Body.Close()

	acceptStatus := map[int]bool{
		http.StatusOK:                   true,
		http.StatusCreated:              true,
		http.StatusNonAuthoritativeInfo: true,
		http.StatusAccepted:             true,
		http.StatusMovedPermanently:     true,
		http.StatusNoContent:            true,
		http.StatusResetContent:         true,
		http.StatusPartialContent:       true,
		http.StatusMultiStatus:          true,
		http.StatusAlreadyReported:      true,
		http.StatusIMUsed:               true,
		http.StatusMultipleChoices:      true,
		http.StatusTemporaryRedirect:    true,
		http.StatusPermanentRedirect:    true,
	}
	if !acceptStatus[res.StatusCode] {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error loading HTTP response body (%w)", err)
	}

	return doc, nil
}
