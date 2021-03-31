package archive

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/romie-gr/romie/internal/utils"
)

var (
	existingFolder = "./testdata"
	extractionPath = "./testdata/archive"
	extractToPath  = "./testdata/new-destination"
	nonWritableDir = "./testdata/non-writable-dir"

	missingArchive = "./testdata/missing.zip"
	nonArchiveFile = "./testdata/archive.txt"

	invalidZip = "./testdata/invalid.zip"
	validZip   = "./testdata/archive.zip"

	invalid7z = "./testdata/invalid.7z"
	valid7z   = "./testdata/archive.7z"

	invalidRar = "./testdata/invalid.rar"
	validRar   = "./testdata/archive.rar"
)

func ExampleExtract() {
	err := Extract("/path/to/archive.zip")
	if err != nil {
		fmt.Printf("Zip archive extraction failure: %v", err)
	} else {
		fmt.Println("Zip archive has been successfully extracted")
	}
	// Output: Zip archive extraction failure: file /path/to/archive.zip not found
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		wantErr bool
	}{
		{
			"Returns error if file does not exist",
			missingArchive,
			true,
		},
		{
			"Returns error if file is a directory",
			existingFolder,
			true,
		},
		{
			"Returns error if file not a valid archive",
			nonArchiveFile,
			true,
		},
		{
			"Zip: Returns error if file not a valid archive",
			invalidZip,
			true,
		},
		{
			"7z: Returns error if file not a valid archive",
			invalid7z,
			true,
		},
		{
			"Rar: Returns error if file not a valid archive",
			invalidRar,
			true,
		},
		{
			"Zip: Returns no error if file is a valid archive",
			validZip,
			false,
		},
		{
			"7z: Returns no error if file is a valid archive",
			valid7z,
			false,
		},
		{
			"Rar: Returns no error if file is a valid archive",
			validRar,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := Extract(tt.source); (err != nil) != tt.wantErr {
				t.Errorf("Extract(%q) error = %v, wantErr %v", tt.source, err, tt.wantErr)
			}

			// Assert file contents and cleanup
			if !tt.wantErr {
				followUpAssertAndCleanUp(t, extractionPath)
			}
		})
	}
}

func ExampleExtractTo() {
	err := ExtractTo("/path/to/archive.zip", "/destination/folder")
	if err != nil {
		fmt.Printf("Zip archive extraction failure: %v", err)
	} else {
		fmt.Println("Zip archive has been successfully extracted")
	}
	// Output: Zip archive extraction failure: file /path/to/archive.zip not found
}

func TestExtractTo(t *testing.T) {
	type args struct {
		source      string
		destination string
	}

	if err := os.Mkdir(nonWritableDir, 0400); err != nil {
		log.Fatalf("Cannot create non writable directory %q", nonWritableDir)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Returns error if file does not exist",
			args{missingArchive, extractToPath},
			true,
		},
		{
			"Returns error if file is a directory",
			args{existingFolder, extractToPath},
			true,
		},
		{
			"Returns error if file not a valid archive",
			args{nonArchiveFile, extractToPath},
			true,
		},
		{
			"Zip: Returns error if file not a valid archive",
			args{invalidZip, extractToPath},
			true,
		},
		{
			"7z: Returns error if file not a valid archive",
			args{invalid7z, extractToPath},
			true,
		},
		{
			"Rar: Returns error if file not a valid archive",
			args{invalidRar, extractToPath},
			true,
		},
		{
			"Returns error if provided path is not writable",
			args{validZip, nonWritableDir},
			true,
		},
		{
			"Zip: Returns no error if file is a valid archive",
			args{validZip, extractToPath},
			false,
		},
		{
			"Rar: Returns no error if file is a valid archive",
			args{validRar, extractToPath},
			false,
		},
		{
			"Rar: Returns no error if file is a valid archive",
			args{validRar, extractToPath},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			skipWindowsNonWritableDirScenario(t, tt.args.destination, tt.name)

			if err := ExtractTo(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("Extract(%q, %q) error = %v, wantErr %v", tt.args.source, tt.args.destination, err, tt.wantErr)
			}
			// Assert file contents and cleanup
			if !tt.wantErr {
				followUpAssertAndCleanUp(t, tt.args.destination)
			}
		})
	}

	// Delete nonWritableDir to ensure test isolation
	_ = os.RemoveAll(nonWritableDir)
}

func skipWindowsNonWritableDirScenario(t *testing.T, destination string, scenarioName string) { // nothelper
	if destination == nonWritableDir && runtime.GOOS == "windows" {
		t.Skipf("Skip %q test in windows", scenarioName)
	}
}

func followUpAssertAndCleanUp(t *testing.T, extractionPath string) { // nothelper
	extractedFile := extractionPath + "/archive.txt"

	// Assert file exists
	if !utils.FolderExists(extractionPath) || !utils.FileExists(extractedFile) {
		t.Errorf("Expected extracted file %q not found", extractedFile)
	}

	// Assert file contents
	content, err := ioutil.ReadFile(extractedFile)
	if err != nil {
		t.Errorf("Getting file content for file %q failed, error = %v", extractedFile, err)
	}

	if expectation := "Text file!\n"; string(content) != expectation {
		t.Errorf("Fail asserting %q content of file %q, matches %q", content, extractedFile, expectation)
	}

	// Cleanup
	_ = os.RemoveAll(extractionPath)
}
