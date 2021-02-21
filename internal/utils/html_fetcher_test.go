package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func ExampleGetHTML() {
	var doc interface{}

	doc, err := GetHTML("https://www.google.com/")
	if err != nil {
		fmt.Println("Document could not be retrieved")
	}

	_, ok := doc.(*goquery.Document)
	if ok {
		fmt.Println("HTML Document successfully returned")
	} else {
		fmt.Println("HTML Document could not be parsed")
	}
	// Output: HTML Document successfully returned
}

func MockServer(testCase string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch testCase {
		case "OK":
			w.WriteHeader(http.StatusOK)
			return
		case "REDIRECT":
			http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)
			return
		case "BADREQUEST":
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		case "NOTFOUND":
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		case "SERVERERROR":
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
			return
		default:
			log.Fatalf("Not a valid test case: %q ", testCase)
			return
		}
	}))

	return ts
}

func TestGetHTML(t *testing.T) {
	tests := []struct {
		name     string
		testCase string
		wantErr  bool
	}{
		{
			"HTTP status code is 200, html object downloaded",
			"OK",
			false,
		},
		{
			"HTTP status code is 301 (redirection), html object downloaded",
			"REDIRECT",
			false,
		},
		{
			"HTTP status code is 400 (bad request), html object not downloaded",
			"BADREQUEST",
			true,
		},
		{
			"HTTP status code is 404 (not found), html object not downloaded",
			"NOTFOUND",
			true,
		},
		{
			"HTTP status code is 500 (server error), html object not downloaded",
			"SERVERERROR",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ts := MockServer(tt.testCase)
			defer ts.Close()
			var doc interface{}
			doc, err := GetHTML(ts.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHTML(%q) error = %v", ts.URL, err)
			} else {
				_, ok := doc.(*goquery.Document)
				if !ok {
					t.Errorf("Returned document from GetHTML(%q) is not a goQuery document", ts.URL)
				}
			}
		})
	}
}
