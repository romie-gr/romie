package utils

import (
	"fmt"
	"testing"
)

var (
	existingFolder    = "./testdata/a-folder-that-exists"
	nonExistingFolder = "./testdata/a-folder-that-does-not-exist"
	existingFile      = "./testdata/a-folder-that-exists/file.txt"
	nonExistingFile   = "./testdata/a-folder-that-exists/missing-file.txt"
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
			existingFolder,
			true,
		},
		{
			"Returns false when given folder does not exist",
			nonExistingFolder,
			false,
		},
		{
			"Returns false when provided path is not a directory",
			existingFile,
			false,
		},
		{
			"Returns false when provided path is empty",
			"",
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := FolderExists(tt.path); got != tt.want {
				t.Errorf("FolderExists(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func ExampleFileExists() {
	exists := FileExists("/missing-file.txt")
	if exists {
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}
	// Output: File does not exist
}

func TestFileExists(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			"Returns true when given file exists",
			existingFile,
			true,
		},
		{
			"Returns false when given file does not exist",
			nonExistingFile,
			false,
		},
		{
			"Returns false when provided path is a directory",
			existingFolder,
			false,
		},
		{
			"Returns false when provided path is empty",
			"",
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.path); got != tt.want {
				t.Errorf("FileExists(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
