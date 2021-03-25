//nolint:noctx
package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

// IsOnline checks the provided URL for connectivity
func IsOnline(url url.URL) error {
	timeout := 2 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url.String())

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	// Non-200 http statuses are considered error
	return fmt.Errorf("timeout or uknown HTTP error, while trying to access %q", url.String())
}

// ParseAndGetDocument fetches the GoQuery document object for a given link
func ParseAndGetDocument(uri string) *goquery.Document {
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
