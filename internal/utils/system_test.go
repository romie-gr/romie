package utils

import (
	"fmt"
	"os"
	"testing"
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
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			"Get environment variable that exists and has a value",
			"KEY_WITH_VALUE",
			"some random value",
			false,
		},
		{
			"Get environment variable that exists but has empty value",
			"EMPTY_KEY",
			"",
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

		if tt.key != "MISSING_KEY" { // should be really missing
			os.Setenv(tt.key, tt.want)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnv(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
