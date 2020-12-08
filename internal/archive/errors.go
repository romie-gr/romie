package archive

import "fmt"

type missingFileError struct {
	FilePath string
}

func (e missingFileError) Error() string {
	return fmt.Sprintf("file %s not found", e.FilePath)
}

type wrongFileExtensionError struct {
	Extension string
}

func (e wrongFileExtensionError) Error() string {
	return fmt.Sprintf("wrong file extension: %s", e.Extension)
}

type openZipError struct {
	Source string
	Err    error
}

func (e openZipError) Error() string {
	return fmt.Sprintf("cannot open zip file: %s: %s", e.Source, e.Err)
}

type extractFileError struct {
	FileName string
	Err      error
}

func (e extractFileError) Error() string {
	return fmt.Sprintf("cannot extract file from archive: %s: %s", e.FileName, e.Err)
}
