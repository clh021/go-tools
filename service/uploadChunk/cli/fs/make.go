package fs

import (
	"os"

	"github.com/spf13/afero"
)

func Mkdir(fs afero.Fs, name string, perm os.FileMode) error {
	return fs.Mkdir(name, perm)
}

func MkdirAll(fs afero.Fs, name string, perm os.FileMode) error {
	return fs.MkdirAll(name, perm)
}

func Create(fs afero.Fs, name string) (afero.File, error) {
	return fs.Create(name)
}
