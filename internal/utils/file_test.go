package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
			if err := CreateFile(tt.filepath); (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	_ = os.Remove(nonExistingFile)
	_ = os.RemoveAll(nonWritableDir)
	_ = os.RemoveAll(nonExistingFolder)
}

func ExamplecreateDir() {
	err := createDir(existingFolder)
	if err != nil {
		fmt.Println("Failed to folder file because it already exists")
	} else {
		fmt.Println("File has been created successfully!")
	}
	// Output: Failed to folder file because it already exists
}

func Test_createDir(t *testing.T) {
	if err := os.Mkdir(nonWritableDir, 0400); err != nil {
		log.Fatalf("Cannot create non writable directory %q", nonWritableDir)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			"Returns nil if directory is created",
			nonExistingFolder,
			false,
		},
		{
			"Returns nil if directory already exists",
			existingFolder,
			false,
		},
		{
			"Returns err if directory cannot be created",
			filepath.Join(nonWritableDir, "newdir"),
			true,
		},
		{
			"Returns err if directory is actually a file",
			existingFile,
			true,
		},
		{
			"Returns err if directory's parent has no read permissions",
			filepath.Join(nonWritableDir, "parent/newdir"),
			true,
		},
		{
			"Returns err when provided path is empty",
			"",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := createDir(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("createDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Delete test data to ensure test isolation
	_ = os.RemoveAll(nonWritableDir)
	_ = os.RemoveAll(nonExistingFolder)
}
