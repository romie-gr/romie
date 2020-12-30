package utils

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// StringContains searches for a string in a series of strings
func StringContains(baseString string, stringlist ...string) bool {
	if strings.TrimSpace(baseString) == "" {
		log.Debug("Input string is empty")
		return false
	}

	for i, s := range stringlist {
		if strings.TrimSpace(s) == "" {
			continue
		}

		if strings.Contains(
			strings.ToLower(baseString),
			strings.ToLower(s),
		) {
			log.Debug(fmt.Sprint("Input string contains argument number ", i, ": ", s))
			return true
		}
	}

	return false
}
