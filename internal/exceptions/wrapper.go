package exceptions

import (
	"fmt"
	"runtime"
)

// getCallerInfo returns file name, line number, function name
// It uses caller to get info about which function caused the called error handling mechanism
// for more info: https://golang.org/pkg/runtime/#Caller
func getCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)

	return file, line, fn.Name()
}

// Wrap allows you to wrap the error message with
// file, line, caller informations which can be useful for reporting.
// Nesting error will wrapped as well.
func Wrap(err error, message string) error {
	f, l, fn := getCallerInfo()
	return fmt.Errorf("file: %v, line: %v, caller: %v message: %s (%w)", f, l, fn, message, err)
}
