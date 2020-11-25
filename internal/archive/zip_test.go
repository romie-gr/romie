package archive

import (
	"archive/zip"
	"bytes"
	"os"
	"testing"

	"github.com/spf13/afero"

	"github.com/romie-gr/romie/internal/utils"
)

func init() {
	utils.AppFS = &afero.Afero{Fs: afero.NewMemMapFs()}
}

var (
	existingFolder = "/folder"
	missingFile    = "/folder/missing.zip"
	nonZipFile     = "/folder/archive.txt"
	invalidZipFile = "/folder/invalid.zip"
	zipArchiveFile = "/folder/archive.zip"
)

func createZipArchive(archivePath string) {
	var buff bytes.Buffer

	// Compress content
	zipW := zip.NewWriter(&buff)
	f, err := zipW.Create("file.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("This is a file"))
	if err != nil {
		panic(err)
	}
	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	// Write output to file
	err = utils.AppFS.WriteFile(archivePath, buff.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestUnzip(t *testing.T) {
	_ = utils.AppFS.Mkdir(existingFolder, 0755)
	_, _ = utils.AppFS.Create(nonZipFile)
	_, _ = utils.AppFS.Create(invalidZipFile)
	createZipArchive(zipArchiveFile)

	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{
			"Returns error if file does not exist",
			missingFile,
			true,
		},
		{
			"Returns error if file is a directory",
			existingFolder,
			true,
		},
		{
			"Returns error if file not a zip archive",
			nonZipFile,
			true,
		},
		{
			"Returns error if file not a valid zip archive",
			invalidZipFile,
			true,
		},
		{
			"Returns no error if file is a valid zip archive",
			zipArchiveFile,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unzip(tt.filepath); (err != nil) != tt.wantErr {
				t.Errorf("Unzip(%q) error = %v, wantErr %v", tt.filepath, err, tt.wantErr)
			}
		})
	}
}
