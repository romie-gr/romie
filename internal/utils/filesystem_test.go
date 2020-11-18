package utils

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

var (
	existingFolder    = "/a-folder-that-exists"
	nonExistingFolder = "/a-folder-that-does-not-exist"
	existingFile      = "/a-folder-that-exists/file.txt"
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
	AppFS = &afero.Afero{Fs: afero.NewMemMapFs()}

	_ = AppFS.Mkdir(existingFolder, 0755)
	_, _ = AppFS.Create(existingFile)

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
				t.Errorf("FolderExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
