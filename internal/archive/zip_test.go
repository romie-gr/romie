package archive

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/romie-gr/romie/internal/utils"
)

var (
	existingFolder = "./testdata"
	missingFile    = "./testdata/missing.zip"
	nonZipFile     = "./testdata/archive.txt"
	invalidZipFile = "./testdata/invalid.zip"
	zipArchiveFile = "./testdata/archive.zip"
	extractionPath = "./testdata/archive"
	extractToPath  = "./testdata/new-destination"
	nonWritableDir = "/sys"
)

func TestUnzip(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		wantErr bool
		cleanUp bool
	}{
		{
			"Returns error if file does not exist",
			missingFile,
			true,
			false,
		},
		{
			"Returns error if file is a directory",
			existingFolder,
			true,
			false,
		},
		{
			"Returns error if file not a zip archive",
			nonZipFile,
			true,
			false,
		},
		{
			"Returns error if file not a valid zip archive",
			invalidZipFile,
			true,
			false,
		},
		{
			"Returns no error if file is a valid zip archive",
			zipArchiveFile,
			false,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := Unzip(tt.source); (err != nil) != tt.wantErr {
				t.Errorf("Unzip(%q) error = %v, wantErr %v", tt.source, err, tt.wantErr)
			}

			// Assert file contents and cleanup
			if tt.cleanUp {
				followUpAssertAndCleanUp(t, extractionPath)
			}
		})
	}
}

func TestUnzipTo(t *testing.T) {
	type args struct {
		source      string
		destination string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		cleanUp bool
	}{
		{
			"Returns error if file does not exist",
			args{missingFile, extractToPath},
			true,
			false,
		},
		{
			"Returns error if file is a directory",
			args{existingFolder, extractToPath},
			true,
			false,
		},
		{
			"Returns error if file not a zip archive",
			args{nonZipFile, extractToPath},
			true,
			false,
		},
		{
			"Returns error if file not a valid zip archive",
			args{invalidZipFile, extractToPath},
			true,
			false,
		},
		{
			"Returns error if provided path is non writable",
			args{zipArchiveFile, nonWritableDir},
			true,
			false,
		},
		{
			"Returns no error if file is a valid zip archive",
			args{zipArchiveFile, extractToPath},
			false,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			skipWindowsNonWritableDirScenario(t, tt.args.destination, tt.name)

			if err := UnzipTo(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("Unzip(%q, %q) error = %v, wantErr %v", tt.args.source, tt.args.destination, err, tt.wantErr)
			}
			// Assert file contents and cleanup
			if tt.cleanUp {
				followUpAssertAndCleanUp(t, tt.args.destination)
			}
		})
	}
}

func skipWindowsNonWritableDirScenario(t *testing.T, destination string, scenarioName string) {
	if destination == nonWritableDir && runtime.GOOS == "windows" {
		t.Skipf("Skip %q test in windows", scenarioName)
	}
}

func followUpAssertAndCleanUp(t *testing.T, extractionPath string) {
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

	expectation := "Text file!\n"
	if string(content) != expectation {
		t.Errorf("Fail asserting %q content of file %q, matches %q", content, extractedFile, expectation)
	}

	// Cleanup
	_ = os.RemoveAll(extractionPath)
}
