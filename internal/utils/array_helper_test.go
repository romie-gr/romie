package utils

import (
	"fmt"
	"testing"
)

func ExampleArrayContains() {
	haystack := []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "India"}
	needle := "Golf"
	index, exists := ArrayContains(haystack, needle)

	if !exists {
		fmt.Printf("Element %s could not be found", needle)
	} else {
		fmt.Printf("Element %s was found with index %d", needle, index)
	}
	// Output: Element Golf was found with index 6
}

func TestArrayContains(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}

	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			"Returns -1 and false if element does not exist",
			args{
				[]string{"Alpha", "Bravo", "Charlie"},
				"Delta",
			},
			-1,
			false,
		},
		{
			"Returns the index and true if element exists",
			args{
				[]string{"Alpha", "Bravo", "Charlie", "Delta", "Echo"},
				"Delta",
			},
			3,
			true,
		},
		{
			"Returns the first index and true if element exists more than once",
			args{
				[]string{"Alpha", "Bravo", "Alpha", "Bravo", "Alpha", "Bravo"},
				"Bravo",
			},
			1,
			true,
		},
		{
			"Returns -1 and false if given slice is empty",
			args{
				[]string{},
				"Alpha",
			},
			-1,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			index, exists := ArrayContains(tt.args.slice, tt.args.val)
			if index != tt.want {
				t.Errorf("ArrayContains() index = %v, want %v", index, tt.want)
			}
			if exists != tt.want1 {
				t.Errorf("ArrayContains() exists = %v, want %v", exists, tt.want1)
			}
		})
	}
}
