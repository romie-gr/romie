package utils

import (
	"os"

	"github.com/romie-gr/romie/internal/exceptions"
)

// EnVar checks if an environment variable exists and returns it.
func EnVar(key string) (string, error) {
	if key == "" {
		return "", exceptions.Wrap(exceptions.ErrArg, "empty argument")
	}

	val, ok := os.LookupEnv(key)

	if val == "" {
		if !ok {
			return "", exceptions.Wrap(exceptions.ErrEnvVar, "variable not found")
		}

		return "", exceptions.Wrap(exceptions.ErrEnvVar, "variable is empty")
	}

	return val, nil
}
