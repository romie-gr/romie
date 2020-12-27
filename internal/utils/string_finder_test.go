package utils

import (
	"fmt"
	"testing"
)

var (
	baseString          = "This is a string"
	baseCaps            = "THIS IS A STRING"
	baseNoSpaces        = "ThisIsAString"
	baseNumeric         = "This 3.14 exists 5236"
	baseNumericNoSpaces = "This3.14Exists5236"
	baseEmpty           = ""
	baseWhitespace      = " 	"
)

func ExampleStringContains() {
	contains := StringContains("This is a string", "not", "are", "simple")
	if contains {
		fmt.Println("String contains at least one of the arguments")
	} else {
		fmt.Println("String does not contain any of the arguments")
	}
	// Output: String does not contain any of the arguments
}

func TestStringContains(t *testing.T) {
	tests := []struct {
		name string
		base string
		args []string
		want bool
	}{
		{"Simple string true search", baseString, []string{"this", "or"}, true},
		{"Simple string false search", baseString, []string{"that", "or"}, false},
		{"String with capitals true search", baseCaps, []string{"this", "or"}, true},
		{"String with capitals false search", baseCaps, []string{"that", "or"}, false},
		{"String without spaces true search", baseNoSpaces, []string{"this", "or"}, true},
		{"String with numbers true search", baseNumeric, []string{"3.1", "236"}, true},
		{"String with numbers without spaces true search", baseNumericNoSpaces, []string{"3.1", "236"}, true},
		{"String with whitespace 1 search", baseWhitespace, []string{"sth"}, false},
		{"String with whitespace 2 search", baseWhitespace, []string{" 	"}, false},
		{"String with whitespace 3 search", baseWhitespace, []string{""}, false},
		{"String empty search", baseEmpty, []string{""}, false},
		{"String full args empty search", baseString, []string{""}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := StringContains(tt.base, tt.args...)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
