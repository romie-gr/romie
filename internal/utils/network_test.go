package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	log "github.com/sirupsen/logrus"
)

func ExampleIsOnline() {
	google, _ := url.Parse("https://google.com")
	hostOK, err := IsOnline(*google)

	if err != nil {
		log.Error(err)
	}

	if hostOK {
		fmt.Println("Host is accessible")
	}
	// Output: Host is accessible
}

func setUpMock(scenario string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch scenario {
		case "OK":
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintln(w, "OK")
			return
		case "Accepted":
			w.WriteHeader(http.StatusAccepted)
			_, _ = fmt.Fprintln(w, "Accepted")
			return
		case "Redirect":
			http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)
			return
		case "BadRequest":
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		case "NotFound":
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		case "RateLimit":
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		case "ServerError":
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
			return
		default:
			log.Fatalf("Unimplemented scenario %q provided", scenario)
			return
		}
	}))

	return ts
}

func Test_IsOnline(t *testing.T) {
	tests := []struct {
		name     string
		scenario string
		hostOk   bool
		wantErr  bool
	}{
		{
			"Succeeds with 200",
			"OK",
			true,
			false,
		},
		{
			"Succeeds with other 200-code (202)",
			"Accepted",
			true,
			false,
		},
		{
			"Succeeds after following a redirect (301)",
			"Redirect",
			true,
			false,
		},
		{
			"Fails if page replies with not generic error (400)",
			"BadRequest",
			false,
			true,
		},
		{
			"Fails if page replies with not found error (404)",
			"NotFound",
			false,
			true,
		},
		{
			"Fails if page replies with too many requests error (429)",
			"RateLimit",
			false,
			true,
		},
		{
			"Fails if page replies with internal server error (500)",
			"ServerError",
			false,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ts := setUpMock(tt.scenario)
			defer ts.Close()
			testURL, _ := url.Parse(ts.URL)

			got, err := IsOnline(*testURL)
			if got != tt.hostOk {
				t.Errorf("IsOnline(%q) got = %v, want %v", testURL.String(), got, tt.hostOk)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("IsOnline(%q) error = %v, wantErr %v", testURL.String(), err, tt.wantErr)
			}
		})
	}
}

func Test_IsOnline_MissingHost(t *testing.T) {
	testURL, _ := url.Parse("https://does-not-exist.romie-gr.romie.com")

	hostOk, err := IsOnline(*testURL)
	if hostOk != false {
		t.Errorf("IsOnline(%q) got = %v, want %v", testURL.String(), hostOk, false)
	}

	if (err != nil) != true {
		t.Errorf("IsOnline(%q) error = %v, wantErr %v", testURL.String(), err, true)
	}
}
