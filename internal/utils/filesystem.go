package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// FolderExists reports whether the provided directory exists.
func FolderExists(path string) bool {
	if path == "" {
		log.Debug("Path is empty")
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug("Path does not exist")
			return false
		}
	}

	return info.IsDir()
}
