package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	existingFolder    = "./testdata/a-folder-that-exists"
	nonExistingFolder = "./testdata/a-folder-that-does-not-exist"
	existingFile      = "./testdata/a-folder-that-exists/file.txt"
	nonExistingFile   = "./testdata/a-folder-that-exists/missing-file.txt"
	nonWritableDir    = "./testdata/non-writable-dir"
)

func createFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

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

func ExampleRemove() {
	filename := "./testdata/a-file-to-be-deleted"
	createFile(filename)
	err := Remove(filename)

	if err == nil {
		fmt.Println("File deleted")
	} else {
		fmt.Println("Unable to delete file")
	}
	// Output: File deleted
}

func TestFileExists(t *testing.T) {
	if err := os.Mkdir(nonWritableDir, 0400); err != nil {
		log.Fatalf("Cannot create non writable directory %q", nonWritableDir)
	}

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
			"Returns false when file is into a folder without read permissions",
			filepath.Join(nonWritableDir, "missing-file.txt"),
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

	_ = os.RemoveAll(nonWritableDir)
}

func TestRemove(t *testing.T) {
	filename := "./testdata/a-file-to-be-deleted"
	createFile(filename)

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			"Delete file that exists",
			filename,
			false,
		},
		{
			"Delete folder that exists",
			existingFolder,
			true,
		},
		{
			"Delete file that does not exist",
			nonExistingFile,
			true,
		},
		{
			"Receive empty path as argument",
			"",
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := Remove(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
