package fs

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/afero"
)

type FsManager struct {
	Fs afero.Fs
}

func NewFsManager(Fs afero.Fs) *FsManager {
	return &FsManager{
		Fs: Fs,
	}
}

func (s *FsManager) GetTree(p string) ([]os.FileInfo, error) {
	return Tree(s.Fs, p)
}

func (s *FsManager) Mkdir(path string) error {
	return Mkdir(s.Fs, path, 0755)
}

func (s *FsManager) MkdirAll(path string) error {
	return MkdirAll(s.Fs, path, 0755)
}

func (s *FsManager) RemoveFile(path string) error {
	return Remove(s.Fs, path)
}

func (s *FsManager) CopyFile(src string, dist string) error {
	return Copy(s.Fs, src, dist)
}

func (s *FsManager) Rename(old string, new_ string) error {
	return Rename(s.Fs, old, new_)
}

func (s *FsManager) Create(name string, fn func(f io.Writer) error) error {
	f, err := Create(s.Fs, name)
	if err != nil {
		return err
	}
	defer f.Close()
	return fn(f)
}

func (s *FsManager) GetFile(filePath string, buffersize int, send func(b []byte) error) error {
	return GetFile(s.Fs, filePath, send)
}

func (s *FsManager) SaveFile(filePath string, fn func(w io.Writer) error) error {
	return SaveFile(s.Fs, filePath, fn)
}

func (s *FsManager) Stat(path string) (os.FileInfo, error) {
	return s.Fs.Stat(path)
}

func (s *FsManager) IsFileExist(path string) bool {
	stat, err := s.Fs.Stat(path)
	if err == nil {
		if stat.IsDir() {
			return false
		}
		return true
	}
	return false
}

func (s *FsManager) IsDirExist(path string) bool {
	ok, err := afero.DirExists(s.Fs, path)
	if !ok || err != nil {
		return false
	}
	return true
}

// Open return http.File
func (s *FsManager) Open(name string) (http.File, error) {
	return s.Fs.Open(name)
}

func (s *FsManager) Zip(dest string, files []string) error {
	return Zip(s.Fs, dest , files)
}