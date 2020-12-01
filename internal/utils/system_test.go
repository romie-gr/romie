package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/romie-gr/romie/internal/exceptions"
)

func ExampleGetEnv() {
	variable := "MY_VAR"
	os.Setenv(variable, "some-value")

	val, err := GetEnv(variable)

	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	// Output: some-value
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name              string
		key               string
		want              string
		throwsError       bool
		wantSpecificError string
	}{
		{
			"Get environment variable that exists and has a value",
			"KEY_WITH_VALUE",
			"some random value",
			false,
			"",
		},
		{
			"Get environment variable that exists but has empty value",
			"EMPTY_KEY",
			"",
			true,
			exceptions.ErrEnvVar.Error(),
		},
		{
			"Get environment variable that does not exist",
			"MISSING_KEY",
			"",
			true,
			exceptions.ErrEnvVar.Error(),
		},
		{
			"Receive empty key as argument",
			"",
			"",
			true,
			exceptions.ErrArg.Error(),
		},
	}

	for _, tt := range tests {
		tt := tt
		os.Setenv(tt.key, tt.want)
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnv(tt.key)

			// Test if error is expected to occur.
			if tt.throwsError && err == nil {
				t.Errorf("GetEnv() error = %v, but it shouldn't through an error!", err)
			}

			// Test for specific error messages, if there is one.
			if err != nil {
				if matches, errMessage := exceptions.Compare(err, tt.wantSpecificError); !matches {
					t.Errorf("GetEnv() error message = %v, want specific error message = %v", errMessage, tt.wantSpecificError)
				}
			}

			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
