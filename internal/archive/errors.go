package archive

import "fmt"

type MissingFileError struct {
	FilePath string
}

func (e MissingFileError) Error() string {
	return fmt.Sprintf("file %s not found", e.FilePath)
}

type WrongFileExtensionError struct {
	Extension string
}

func (e WrongFileExtensionError) Error() string {
	return fmt.Sprintf("wrong file extension: %s", e.Extension)
}

type CreateDirectoryError struct {
	Err error
}

func (e CreateDirectoryError) Error() string {
	return fmt.Sprintf("cannot create directory: %s", e.Err)
}

type OpenZipError struct {
	FilePath string
	Err      error
}

func (e OpenZipError) Error() string {
	return fmt.Sprintf("cannot open zip file: %s: %s", e.FilePath, e.Err)
}

type ExtractFileError struct {
	FileName string
	Err      error
}

func (e ExtractFileError) Error() string {
	return fmt.Sprintf("cannot extract file from archive: %s: %s", e.FileName, e.Err)
}

type OpenFileError struct {
	Location string
	Write    bool
	Err      error
}

func (e OpenFileError) Error() string {
	if e.Write {
		return fmt.Sprintf("cannot open file for writing: %s: %s", e.Location, e.Err)
	}

	return fmt.Sprintf("cannot open file: %s: %s", e.Location, e.Err)
}

type StatFileError struct {
	Location string
	Err      error
}

func (e StatFileError) Error() string {
	return fmt.Sprintf("cannot retrieve fileinfo for file: %s: %s", e.Location, e.Err)
}

type WriteFileError struct {
	Location string
	Err      error
}

func (e WriteFileError) Error() string {
	return fmt.Sprintf("cannot write to file: %s: %s", e.Location, e.Err)
}
