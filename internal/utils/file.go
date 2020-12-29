package utils

import (
	"fmt"
	"os"

	osfile "path/filepath"
)

// CreateFile creates a file on the given filepath
// along with any necessary parents, and returns nil,
// or else returns an error.
func CreateFile(filepath string) error {
	// Avoid truncate, but return error instead
	if FileExists(filepath) {
		return fmt.Errorf("file %s already exists", filepath)
	}

	// Create the parent directory if doesn't exist
	if directory := osfile.Dir(filepath); !FolderExists(directory) {
		if err := createDir(directory); err != nil {
			return fmt.Errorf("failed to create directory %v", directory)
		}
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %w", filepath, err)
	}
	defer file.Close()

	return nil
}

func createDir(newpath string) error {
	return os.MkdirAll(newpath, os.ModePerm)
}
