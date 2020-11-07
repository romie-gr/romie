package utils

import (
	"fmt"
	"testing"
)

func ExampleFolderExists() {
	exists := FolderExists("/a-non-existing-folder")
	if exists {
		fmt.Println("Folder exists")
	} else {
		fmt.Println("Folder does not exist")
	}
	// Output: Folder does not exist
}

func TestFolderExists(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			"Returns true when given folder exists",
			"./testdata/a-folder-that-exists",
			true,
		},
		{
			"Returns false when given folder does not exist",
			"./testdata/a-folder-that-does-not-exist",
			false,
		},
		{
			"Returns false when provided path is not a directory",
			"./testdata/file.txt",
			false,
		},
		{
			"Returns false when provided path is empty",
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FolderExists(tt.path); got != tt.want {
				t.Errorf("FolderExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
