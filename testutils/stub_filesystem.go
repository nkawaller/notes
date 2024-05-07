package testutils

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"os"
	"testing/fstest"
)

type StubFileSystem struct {
	FS               fstest.MapFS
	ContentRoot      string
	TemplateLocation string
}

// Create creates a file with the given name
func (s StubFileSystem) Create(name string) (io.Writer, error) {
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	s.FS[name] = &fstest.MapFile{
		Data: buf.Bytes(),
	}
	return writer, nil
}

func (s StubFileSystem) MkdirAll(path string, perm os.FileMode) error {
	dir := fstest.MapFile{Mode: perm | os.ModeDir}
	s.FS[path] = &dir // MapFs needs a pointer
	return nil
}

func (s StubFileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return s.FS.ReadDir(name)
}

func (s StubFileSystem) ReadFile(filename string) ([]byte, error) {
	return s.FS.ReadFile(filename)
}

func (s StubFileSystem) Stat(name string) (os.FileInfo, error) {
	return s.FS.Stat(name)
}

func (s StubFileSystem) GetContentRoot() string {
	return s.ContentRoot
}

func (s StubFileSystem) GetTemplateLocation() string {
	return s.TemplateLocation
}
