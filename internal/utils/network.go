//nolint:noctx
package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
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
