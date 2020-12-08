package archive

import (
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-unarr"
	log "github.com/sirupsen/logrus"

	"github.com/romie-gr/romie/internal/utils"
)

func getFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Unzip Decompresses an archive inside its containing directory
func Unzip(source string) error {
	// Extract file in current directory
	extractPath := getFileNameWithoutExtension(source)
	return extract(source, extractPath)
}

// UnzipTo Decompresses an archive inside the given directory
func UnzipTo(source string, destination string) error {
	return extract(source, destination)
}

func extract(source string, destination string) error {
	// Check for file existence
	if !utils.FileExists(source) {
		log.Error("Error trying to identify path")
		return missingFileError{source}
	}

	// Check file has .zip extension
	if ext := filepath.Ext(source); ext != ".zip" {
		log.Errorf("Invalid archive extension '%s' for zip method", ext)
		return wrongFileExtensionError{ext}
	}

	// Check file is archive
	archive, err := unarr.NewArchive(source)
	if err != nil {
		log.Error("Archive is not valid")
		return openZipError{source, err}
	}
	defer archive.Close()

	// Extract file in current directory
	_, err = archive.Extract(destination)
	if err != nil {
		log.Error(err.Error())
		return extractFileError{source, err}
	}

	return nil
}
