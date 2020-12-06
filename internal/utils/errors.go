package utils

import "fmt"

// ArgError denotes that a function's argument was passed incorrectly.
type ArgError struct {
	msg string
}

func (e ArgError) Error() string {
	return fmt.Sprintf("Argument error: %s", e.msg)
}

// EnvVarError tracks environment variable errors.
type EnvVarError struct {
	key string
	msg string
}

func (e EnvVarError) Error() string {
	return fmt.Sprintf("Env variable error: %s: %s", e.msg, e.key)
}
