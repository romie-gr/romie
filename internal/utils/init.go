package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// AppFS Application filesystem.
var AppFS = &afero.Afero{Fs: afero.NewOsFs()}

func init() {
	// Set debug mode on
	if os.Getenv("ROMIE_DEBUG") == "on" {
		// See https://github.com/romie-gr/romie/issues/154
		log.SetLevel(log.DebugLevel)
	}
}
