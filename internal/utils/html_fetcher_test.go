package utils

import (
	"fmt"
	"io"
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
		case "a":
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "<html><body>Hello World!</body></html>")
			return
		case "b":
			http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)
			io.WriteString(w, "<html><body>Hello World!</body></html>")
			return
		case "c":
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		case "d":
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		case "e":
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
			"OK",
			"a",
			false,
		},
		{
			"Redirect",
			"b",
			false,
		},
		{
			"BadRequest",
			"c",
			true,
		},
		{
			"NotFound",
			"d",
			true,
		},
		{
			"ServerError",
			"d",
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
