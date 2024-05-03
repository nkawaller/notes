package testutils

import (
	"bytes"
	"os"
	"testing/fstest"
)

type StubFileSystem struct {
	FS               fstest.MapFS
	ContentRoot      string
	TemplateLocation string
}

func (s StubFileSystem) Create(name string) (*os.File, error) {
    // Create a buffer to write the content
    buffer := &bytes.Buffer{}

    // Create a new fs.File instance with the buffer as its data
    file := &fstest.MapFile{
        Data: buffer.Bytes(),
    }

    // Add the file to the MapFS
    s.FS[name] = file

    // Return nil for the os.File
    return nil, nil
}

func (s StubFileSystem) MkdirAll(path string, perm os.FileMode) error {
	dir := fstest.MapFile{Mode: perm | os.ModeDir}
	s.FS[path] = &dir // MapFs needs a pointer
	return nil
}

func (s StubFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
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
