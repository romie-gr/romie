package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func ExampleCreateFile() {
	err := CreateFile(existingFile)
	if err != nil {
		fmt.Println("Failed to create file because it already exists")
	} else {
		fmt.Println("File has been created successfully!")
	}
	// Output: Failed to create file because it already exists
}

func TestCreateFile(t *testing.T) {
	if err := os.Mkdir(nonWritableDir, 0400); err != nil {
		log.Fatalf("Cannot create non writable directory %q", nonWritableDir)
	}

	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{
			"Returns nil if file is created",
			nonExistingFile,
			false,
		},
		{
			"Returns err if file is exists",
			existingFile,
			true,
		},
		{
			"Returns nil if file is created along with parent dirs",
			filepath.Join(nonExistingFolder, "file.txt"),
			false,
		},
		{
			"Returns err if cannot write the file",
			filepath.Join(nonWritableDir, "should-not-write-this"),
			true,
		},
		{
			"Returns err when provided filepath is empty",
			"",
			true,
		},
		{
			"Returns err if parent directory couldn't be created",
			filepath.Join(nonWritableDir, "parent/newdir/file.txt"),
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			skipWindowsNonWritableDirScenario(t, tt.filepath, tt.name)

			if err := CreateFile(tt.filepath); (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	_ = os.Remove(nonExistingFile)
	_ = os.RemoveAll(nonWritableDir)
	_ = os.RemoveAll(nonExistingFolder)
}

func skipWindowsNonWritableDirScenario(t *testing.T, file string, scenarioName string) {
	if strings.Contains(filepath.Base(nonWritableDir), file) && runtime.GOOS == "windows" {
		t.Skipf("Skip %q test in windows", scenarioName)
	}
}
