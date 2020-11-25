package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/romie-gr/romie/internal/utils"
)

// Unzip Decompresses an archive inside its containing directory
func Unzip(filePath string) error {
	// Check for file existence
	if !utils.FileExists(filePath) {
		log.Error("Error trying to identify path")
		return MissingFileError{filePath}
	}

	// Check file has .zip extension
	if ext := filepath.Ext(filePath); ext != ".zip" {
		log.Error(fmt.Sprintf("Invalid archive extension '%s' for zip method", ext))
		return WrongFileExtensionError{ext}
	}

	// Extract file in current directory
	extractPath := getFileNameWithoutExtension(filePath)
	return extract(filePath, extractPath)
}

func getFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func extract(source string, destination string) error {
	log.Infof("extracting %s into %s", source, destination)
	err := utils.AppFS.MkdirAll(destination, 0755)
	if err != nil {
		return CreateDirectoryError{err}
	}

	file, err := utils.AppFS.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	reader, err := zip.NewReader(file, fileStat.Size())
	if err != nil {
		return OpenZipError{source, err}
	}

	for _, file := range reader.File {
		err := unzipFile(destination, file)
		if err != nil {
			return ExtractFileError{file.Name, err}
		}
	}

	log.Debugf("extract of %s was successful", source)
	return nil
}

func unzipFile(destination string, file *zip.File) error {
	contents, err := file.Open()
	if err != nil {
		return ExtractFileError{file.Name, err}
	}
	defer contents.Close()

	if file.FileInfo().IsDir() {
		return nil
	}

	savedLocation := path.Join(destination, file.Name)
	directory := path.Dir(savedLocation)
	err = utils.AppFS.MkdirAll(directory, 0755)
	if err != nil {
		return CreateDirectoryError{err}
	}

	mode := file.Mode()
	newFile, err := utils.AppFS.OpenFile(savedLocation, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return OpenFileError{savedLocation, err}
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, contents)
	if err != nil {
		return WriteFileError{savedLocation, err}
	}

	return nil
}
