package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

var (
	AppFs     afero.Fs
	FSUtil *afero.Afero
)

func init() {
	AppFs = afero.NewOsFs()
	FSUtil = &afero.Afero{Fs: AppFs}
}

// FolderExists reports whether the provided directory exists.
func FolderExists(path string) bool {
	if path == "" {
		log.Debug("Path is empty")
		return false
	}

	exists, err := FSUtil.DirExists(path)
	if err != nil {
		log.Debug(fmt.Printf("Error: %s", err.Error()))
		return false
	}

	return exists
}
