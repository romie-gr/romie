package utils

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// TimeTrack starts a deferred timer, in order to profile the execution of a function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("function %s took %s", name, elapsed)
}
