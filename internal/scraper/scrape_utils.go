package scraper

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func LogAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Info(logStr)
		return nil
	}
}
