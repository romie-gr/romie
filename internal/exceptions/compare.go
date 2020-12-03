package exceptions

import (
	"strings"
)

func isWrapped(err error) bool {
	hasFile := strings.Contains(err.Error(), "file:")
	hasCaller := strings.Contains(err.Error(), "caller:")
	hasMessage := strings.Contains(err.Error(), "message:")

	if hasFile && hasCaller && hasMessage {
		return true
	}

	return false
}

// CompareWrappedErrorMessages it's useful for testing specific Wrapped error messages.
func compareWrappedErrorMessages(err error, wantSpecificError string) (bool, string) {
	errMessage := strings.SplitAfter(err.Error(), "message: ")[1]

	if strings.Contains(errMessage, wantSpecificError) {
		return true, errMessage
	}

	return false, errMessage
}

// compareErrorMessages compares an error message with a string.
func compareErrorMessages(err error, wantSpecificError string) (bool, string) {
	if strings.Contains(err.Error(), wantSpecificError) {
		return true, wantSpecificError
	}

	return false, wantSpecificError
}

// Compare it's useful for testing specific error messages.
func Compare(err error, wantSpecificError string) (match bool, message string) {
	if isWrapped(err) {
		match, message = compareWrappedErrorMessages(err, wantSpecificError)
	} else {
		match, message = compareErrorMessages(err, wantSpecificError)
	}

	return match, message
}
