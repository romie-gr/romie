package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// GetEnv returns the value of the environment variable specified by key, if the variable is set.
func GetEnv(key string) (string, error) {
	if key == "" {
		log.Debug("key is emty")
		return "", ArgError{"empty argument"}
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		log.Error("Error while getting environment variable")
		return "", EnvVarError{key, "environment variable not found"}
	}

	if val == "" {
		return "", EnvVarError{key, "environment variable is empty"}
	}

	return val, nil
}
