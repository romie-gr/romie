package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Exists reports whether the named file or directory exists.
func exists(path string, isDir bool) bool {
	if path == "" {
		log.Debug("Path is empty")
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) || os.IsPermission(err) {
			return false
		}
	}

	return isDir == info.IsDir()
}

// FolderExists reports whether the provided directory exists.
func FolderExists(path string) bool {
	return exists(path, true)
}

// FileExists reports whether the provided file exists.
func FileExists(path string) bool {
	return exists(path, false)
}

// Remove deletes the specified file
func Remove(path string) error {
	if path == "" {
		log.Debug("Path is empty")
		return argError{"empty argument"}
	}

	if !FileExists(path) {
		return missingFileError{path}
	}

	return os.Remove(path)
}
