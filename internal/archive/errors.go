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
	return fmt.Sprintf("invalid or unsupported file extension: %s", e.Extension)
}

type openArchiveError struct {
	Source string
	Err    error
}

func (e openArchiveError) Error() string {
	return fmt.Sprintf("cannot open archive: %s: %s", e.Source, e.Err)
}

type extractFileError struct {
	FileName string
	Err      error
}

func (e extractFileError) Error() string {
	return fmt.Sprintf("cannot extract file from archive: %s: %s", e.FileName, e.Err)
}
