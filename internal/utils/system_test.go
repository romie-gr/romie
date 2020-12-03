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

// Prepate testcase parameters.
type testcase struct {
	name              string
	key               string
	want              string
	wantSpecificError string
}

func isEmpty(envVariable string) bool {
	return envVariable == ""
}

func runErrorExpectedTests(err error, tt testcase) (testResult error) {
	// When error is expected to occur, but it doesn't.
	if err == nil {
		testResult = fmt.Errorf("GetEnv() = %w, but it shouldn't through an error", err)
	}
	// When error is expected and occurs, but for the wrong reason.
	if matches, errMessage := exceptions.Compare(err, tt.wantSpecificError); !matches {
		testResult = fmt.Errorf("GetEnv() = %v, want %v", errMessage, tt.wantSpecificError)
	}

	return testResult
}

func runNoErrorExpectedTests(got string, err error, tt testcase) (testResult error) {
	// When error is not expected to occur, but it does.
	if err != nil {
		testResult = fmt.Errorf("GetEnv() = %w, but it shouldn't through an error", err)
	}
	// When there's no error, but the value of the env variable is wrong
	if got != tt.want {
		testResult = fmt.Errorf("GetEnv() = %v, want %v", got, tt.want)
	}

	return testResult
}

func testScenario(got string, err error, tt testcase) (testResult error) {
	envVariableValue := tt.want

	if isEmpty(envVariableValue) {
		testResult = runErrorExpectedTests(err, tt)
	} else {
		testResult = runNoErrorExpectedTests(got, err, tt)
	}

	return testResult
}

func TestGetEnv(t *testing.T) {
	// Define testing scenarios
	testsuite := []testcase{
		{
			"Get environment variable that exists and has a value",
			"KEY_WITH_VALUE",
			"some random value",
			"",
		},
		{
			"Get environment variable that exists but has empty value",
			"EMPTY_KEY",
			"",
			exceptions.ErrEnvVar.Error(),
		},
		{
			"Get environment variable that does not exist",
			"MISSING_KEY",
			"",
			exceptions.ErrEnvVar.Error(),
		},
		{
			"Receive empty key as argument",
			"",
			"",
			exceptions.ErrArg.Error(),
		},
	}

	for _, tt := range testsuite {
		tt := tt

		if tt.key != "MISSING_KEY" { // should be really missing
			os.Setenv(tt.key, tt.want)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnv(tt.key)
			if result := testScenario(got, err, tt); result != nil {
				t.Error(result)
			}
		})
	}
}
