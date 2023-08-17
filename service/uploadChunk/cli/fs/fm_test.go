package fs

import (
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGetTree(t *testing.T) {
	manager := NewFsManager()
	infos, err := manager.GetTree("/tmp")
	assert.Nil(t, err)
	assert.True(t, len(infos) > 0)
}

func TestCreate(t *testing.T) {
	manager := NewFsManager()

	err := manager.Create("/tmp/nstd-test.txt", func(f io.Writer) error {
		_, err := f.Write([]byte("nanoapp-nstd test string."))
		return err
	})
	assert.Nil(t, err)
	_, err = os.Stat("/tmp/nstd-test.txt")
	assert.Nil(t, err)
}

func TestSaveFile(t *testing.T) {
	manager := NewFsManager()

	f, err := os.Open("/tmp/nstd-test.txt")
	assert.Nil(t, err)

	err = manager.SaveFile("/tmp/nstd-test-1.txt", func(w io.Writer) error {
		_, err := io.Copy(w, f)
		return err
	})
	assert.Nil(t, err)
}

func TestGetFile(t *testing.T) {
	manager := NewFsManager()

	err := manager.GetFile("/tmp/nstd-test.txt", 0, func(b []byte) error {
		assert.Equal(t, []byte("nanoapp-nstd test string."), b)
		return nil
	})
	assert.Nil(t, err)

	err = manager.GetFile("/tmp/nstd-test-1.txt", 0, func(b []byte) error {
		assert.Equal(t, []byte("nanoapp-nstd test string."), b)
		return nil
	})
	assert.Nil(t, err)
}

func TestMkdir(t *testing.T) {
	manager := NewFsManager()
	os.Remove("/tmp/nstd-test-dir")
	err := manager.Mkdir("/tmp/nstd-test-dir")
	assert.Nil(t, err)
}

func TestCopy(t *testing.T) {
	manager := NewFsManager()
	err := manager.CopyFile("/tmp/nstd-test.txt", "/tmp/nstd-test-backup.txt")
	assert.Nil(t, err)

	err = manager.CopyFile("/tmp/nstd-test-dir", "/tmp/nstd-test-dir-2")
	assert.Nil(t, err)

	err = manager.CopyFile("/tmp/nstd-test.txt", "/tmp/nstd-test-dir-2")
	assert.Nil(t, err)

	// Test Copy Link
	os.Remove("/tmp/nstd-test-link.txt")
	err = os.Symlink("/tmp/nstd-test.txt", "/tmp/nstd-test-link.txt")
	assert.Nil(t, err)
	link1, err := os.Lstat("/tmp/nstd-test-link.txt")
	assert.Nil(t, err)
	assert.Equal(t, os.ModeSymlink, link1.Mode()&os.ModeSymlink)

	err = manager.CopyFile("/tmp/nstd-test-link.txt", "/tmp/nstd-test-link-2.txt")
	assert.Nil(t, err)

	f, err := os.Lstat("/tmp/nstd-test-link-2.txt")
	assert.Nil(t, err)
	assert.Equal(t, os.ModeSymlink, f.Mode()&os.ModeSymlink)
}

func TestZip(t *testing.T) {
	osfs := afero.NewOsFs()
	os.Remove("/tmp/nstd-test.zip")
	err := Zip(osfs, "/tmp/nstd-test.zip", []string{
		"/tmp/nstd-test-link.txt",
		"/tmp/nstd-test.txt",
		"/tmp/nstd-test-dir-2",
	})
	assert.Nil(t, err)
}

func TestUnzip(t *testing.T) {
	osfs := afero.NewOsFs()
	err := Unzip(osfs, "/tmp/nstd-test.zip", "/tmp/nstd-test-unzip")
	assert.Nil(t, err)

	ok, err := afero.Exists(osfs, "/tmp/nstd-test-unzip/nstd-test-link.txt")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = afero.Exists(osfs, "/tmp/nstd-test-unzip/nstd-test.txt")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = afero.DirExists(osfs, "/tmp/nstd-test-unzip/nstd-test-dir-2")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = afero.Exists(osfs, "/tmp/nstd-test-unzip/nstd-test-dir-2/nstd-test.txt")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemove(t *testing.T) {
	manager := NewFsManager()

	err := manager.RemoveFile("/tmp/nstd-test-link.txt")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test-link-2.txt")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test.txt")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test-1.txt")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test-backup.txt")
	assert.Nil(t, err)

	// remove dir
	err = manager.RemoveFile("/tmp/nstd-test-dir")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test-dir-2")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test-unzip")
	assert.Nil(t, err)

	err = manager.RemoveFile("/tmp/nstd-test.zip")
	assert.Nil(t, err)
}

func TestRename(t *testing.T) {
	manager := NewFsManager()

	TestCreate(t)

	err := manager.Rename("/tmp/nstd-test.txt", "/tmp/nstd-test-rename.txt")
	assert.Nil(t, err)
	defer os.Remove("/tmp/nstd-test-rename.txt")

	TestMkdir(t)
	err = manager.Rename("/tmp/nstd-test-dir", "/tmp/nstd-test-rename-dir")
	assert.Nil(t, err)
	f, err := os.Stat("/tmp/nstd-test-rename-dir")
	assert.Nil(t, err)
	assert.True(t, f.IsDir())
	defer os.Remove("/tmp/nstd-test-rename-dir")
}

func TestLink(t *testing.T) {
	// manager := NewFsManager()
	// s, _ := manager.Fs.ReadlinkIfPossible("/home/ubuntu/haha")
	// fmt.Println(s)
}

func TestConvertAbsoluteDir(t *testing.T) {
	exceptHome, err := os.UserHomeDir()
	assert.Nil(t, err)

	home, err := ConvertAbsoluteDir("HOME")
	assert.Nil(t, err)
	assert.Equal(t, exceptHome, home)
}
