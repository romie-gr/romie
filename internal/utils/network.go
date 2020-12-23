//nolint:noctx
package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func IsOnline(url url.URL) (bool, error) {
	timeout := 2 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url.String())

	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	// Non-200 http statuses are considered error
	err = fmt.Errorf("timeout or uknown HTTP error, while trying to access %q", url.String())

	return false, err
}
