package fs

import (
	"github.com/spf13/afero"
)

func Remove(fs afero.Fs, path string) error {
	f, err := fs.Stat(path)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return fs.RemoveAll(path)
	}
	return fs.Remove(path)
}

func Rename(fs afero.Fs, oldName string, newName string) error {
	_, err := fs.Stat(oldName)
	if err != nil {
		return err
	}
	return fs.Rename(oldName, newName)
}
