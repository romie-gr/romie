package utils

import (
	"fmt"
	"os"
	"testing"
)

var (
	keyWithValue = "KEY_WITH_VALUE"
	emptyKey     = "EMPTY_KEY"
)

var varValue = map[string]string{
	keyWithValue: "value",
	emptyKey:     "",
}

func setupTest() {
	for key, value := range varValue {
		os.Setenv(key, value)
	}
}

func ExampleEnVar() {
	variable := "MY_VAR"
	os.Setenv(variable, "some-value")

	val, err := EnVar(variable)

	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	// Output: some-value
}

func TestEnVar(t *testing.T) {
	setupTest()

	tests := []struct {
		name        string
		key         string
		want        string
		throwsError bool
	}{
		{
			"Get environment variable that exists and has a value",
			keyWithValue,
			varValue[keyWithValue],
			false,
		},
		{
			"Get environment variable that exists but has empty value",
			emptyKey,
			varValue[emptyKey],
			true,
		},
		{
			"Get environment variable that does not exist",
			"MISSING_KEY",
			"",
			true,
		},
		{
			"Receive empty key as argument",
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := EnVar(tt.key)
			if got != tt.want {
				t.Errorf("EnVar() = %v, want %v", got, tt.want)
			}
			gotError := (err != nil)
			if gotError != tt.throwsError {
				t.Errorf("EnVar() error = %v, but throws error = %v", err, tt.throwsError)
				return
			}
		})
	}
}
