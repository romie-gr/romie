package utils_test

import (
	"fmt"
	"testing"

	"github.com/romie-gr/romie/internal/utils"

	"github.com/spf13/afero"
)

func init() {
	utils.AppFs = afero.NewMemMapFs()
	utils.FSUtil = &afero.Afero{Fs: utils.AppFs}

	_ = utils.FSUtil.Mkdir("/a-folder-that-exists", 0755)
	_, _ = utils.FSUtil.Create("/a-folder-that-exists/file.txt")
}

func ExampleFolderExists() {
	exists := utils.FolderExists("/a-non-existing-folder")
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
			"/a-folder-that-exists",
			true,
		},
		{
			"Returns false when given folder does not exist",
			"/a-folder-that-does-not-exist",
			false,
		},
		{
			"Returns false when provided path is not a directory",
			"/a-folder-that-exists/file.txt",
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
			if got := utils.FolderExists(tt.path); got != tt.want {
				t.Errorf("FolderExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
