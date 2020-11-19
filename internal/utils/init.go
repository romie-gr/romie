package utils

import "github.com/spf13/afero"

// AppFS Application filesystem.
var AppFS = &afero.Afero{Fs: afero.NewOsFs()}
