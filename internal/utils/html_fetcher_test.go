package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func ExampleGetHTMLObj() {
	doc, err := GetHTMLObj("https://www.google.com/")
	if err != nil {
		fmt.Println("Document could not be retrieved")
	_, ok = doc.(goquery.Document)
	if ok {
		fmt.Println("HTML Document successfully returned")
	} else {
		fmt.Println("HTML Document could not be parsed")
	}
	// Output: HTML Document successfully returned
}