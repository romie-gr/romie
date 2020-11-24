package exceptions

import "errors"

// New converts string to error.
func New(err string) error {
	return errors.New(err)
}

// ErrArg denotes that a function's argument was passed incorrectly.
var ErrArg = errors.New("argument error")

// ErrEnvVar is caused by environment variable exceptions.
var ErrEnvVar = errors.New("environment Variable error")
