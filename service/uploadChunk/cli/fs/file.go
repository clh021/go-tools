package fs

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

const (
	// NOTE: if the BUFFERSIZE too large, will cause
	// `Response closed without grpc-status (Headers only)` error on JavaScript.
	BUFFERSIZE = 1024 * 64
)

func GetFile(fs afero.Fs, name string, send func(b []byte) error) error {
	f, err := fs.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			sendErr := send(buf[:n])
			if sendErr != nil {
				return sendErr
			}
			if err == io.EOF {
				return nil
			}
		}

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

	}
	return nil
}

func SaveFile(fs afero.Fs, name string, fn func(w io.Writer) error) error {
	f, err := fs.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return fn(f)
}

func Tree(fs afero.Fs, name string) ([]os.FileInfo, error) {
	res := make([]os.FileInfo, 0)
	stat, err := fs.Stat(name)
	if err != nil {
		return res, err
	}

	if !stat.IsDir() {
		return res, fmt.Errorf("%s is a file,not directory", name)
	}

	return afero.ReadDir(fs, name)
}

func Zip(fs afero.Fs, dest string, files []string) error {
	if filepath.Ext(dest) != ".zip" {
		return fmt.Errorf("Destination file don't zip file. Please use zip extension")
	}

	exist, err := afero.Exists(fs, dest)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("Destination file already exist.")
	}

	f, err := fs.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	base := filepath.Dir(dest)

	w := zip.NewWriter(f)
	for _, v := range files {
		stat, err := fs.Stat(v)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			err = zipDir(fs, base, v, w)
		} else {
			err = zipFile(fs, base, v, w)
		}
		if err != nil {
			return err
		}
	}
	return w.Close()
}

func zipDir(fs afero.Fs, dest string, name string, w *zip.Writer) error {
	dir, _ := fs.Open(name)
	obs, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	var errs []error
	for _, obj := range obs {
		objName := filepath.Join(name, obj.Name())
		if obj.IsDir() {
			err = zipDir(fs, dest, objName, w)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			err = zipFile(fs, dest, objName, w)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}

	if errString != "" {
		return errors.New(errString)
	}
	return nil
}

func zipFile(fs afero.Fs, dest, name string, w *zip.Writer) error {
	b, err := afero.ReadFile(fs, name)
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(dest, name)
	if err != nil {
		return err
	}

	zf, err := w.Create(rel)
	if err != nil {
		return err
	}

	_, err = zf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func Unzip(fs afero.Fs, name string, to string) error {
	f, err := fs.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	exist, err := afero.DirExists(fs, to)
	if err != nil {
		return err
	}
	if !exist {
		err = fs.MkdirAll(to, 0755)
		if err != nil {
			return err
		}
	}

	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return err
	}

	for _, v := range r.File {
		err := unzipTo(fs, v, to)
		if err != nil {
			return err
		}
	}
	return nil
}

func unzipTo(fs afero.Fs, file *zip.File, to string) error {
	if filepath.Dir(file.Name) == filepath.Clean(file.Name) {
		return nil
	}

	name := filepath.Join(to, file.Name)
	dir := filepath.Dir(name)
	if dir != "." {
		mkerr := fs.MkdirAll(dir, 0755)
		if mkerr != nil {
			return mkerr
		}
	}

	e, err := afero.Exists(fs, name)
	if err != nil {
		return err
	}
	if e {
		return nil
	}

	fw, err := fs.Create(name)
	if err != nil {
		return err
	}
	defer fw.Close()

	fr, err := file.Open()
	if err != nil {
		return err
	}
	defer fr.Close()

	_, err = io.Copy(fw, fr)
	return err
}

// LinkTo find the absolute path of link.
// func LinkTo(linker afero.LinkReader, name string) (string, error) {
// 	// return linker.ReadlinkIfPossible(name)
// 	return "TODO", nil
// }

func IsDir(fs afero.Fs, name string) (bool, error) {
	return afero.IsDir(fs, name)
}

// func RelativePath(fs afero.Fs, name string) (string, error) {
// 	panic("")
// }
