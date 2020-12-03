package utils

import (
	"os"

	"github.com/romie-gr/romie/internal/exceptions"
)

// GetEnv returns value of the environment variable named by the key if it's populated.
func GetEnv(key string) (string, error) {
	if key == "" {
		return "", exceptions.Wrap(exceptions.ErrArg, "empty argument")
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		return "", exceptions.Wrap(exceptions.ErrEnvVar, "environment variable not found")
	}

	if val == "" {
		return "", exceptions.Wrap(exceptions.ErrEnvVar, "environment variable is empty")
	}

	return val, nil
}
