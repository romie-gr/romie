package utils

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// StringContains searches for a string in a series of strings
func StringContains(baseString string, stringlist ...string) bool {
	if strings.TrimSpace(baseString) == "" {
		log.Debug("String is empty")
		return false
	}
	for index, str := range stringlist {
		if strings.TrimSpace(str) == "" {
			continue
		}
		if strings.Contains(
			strings.ToLower(baseString),
			strings.ToLower(str),
		) {
			log.Debug(fmt.Sprint("String contains arguement number ", index, ": ", str))
			return true
		}
	}
	log.Debug("String does not contain arguements")
	return false
}
