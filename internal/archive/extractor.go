package archive

import (
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-unarr"
	"github.com/romie-gr/romie/internal/utils"
)

type extractor struct {
	source      string
	destination string
}

// Extract Decompresses an archive inside its containing directory
func Extract(source string) error {
	// extract file in current directory
	destination := getFileNameWithoutExtension(source)

	extractor := extractor{
		source:      source,
		destination: destination,
	}

	return extractor.extract()
}

// ExtractTo Decompresses an archive inside the given directory
func ExtractTo(source string, destination string) error {
	extractor := extractor{
		source:      source,
		destination: destination,
	}

	return extractor.extract()
}

func (e extractor) extract() error {
	// Check for file existence
	if !utils.FileExists(e.source) {
		return missingFileError{e.source}
	}

	// Check file extension is one of the allowed archive extensions (.zip, .rar, or .7z)
	if ext := filepath.Ext(e.source); !hasValidExtension(ext) {
		return wrongFileExtensionError{ext}
	}

	// Check file is archive
	archive, err := unarr.NewArchive(e.source)
	if err != nil {
		return openArchiveError{e.source, err}
	}
	defer archive.Close()

	// Extract file in provided directory
	_, err = archive.Extract(e.destination)
	if err != nil {
		return extractFileError{e.source, err}
	}

	return nil
}

func hasValidExtension(extension string) bool {
	switch extension {
	case
		".zip",
		".rar",
		".7z":
		return true
	}

	return false
}

func getFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
