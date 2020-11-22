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
		// TODO: This is temporary. Remove it once Cobra/Viper is configured, to handle it as a boolean flags
		log.SetLevel(log.DebugLevel)
	}
}
