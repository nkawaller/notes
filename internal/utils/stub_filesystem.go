package utils

import (
    "os"
    "testing/fstest"
)

type StubFileSystem struct {
	FS               fstest.MapFS
	ContentRoot      string
	TemplateLocation string
}

func (s StubFileSystem) MkdirAll(path string, perm os.FileMode) error {
	dir := fstest.MapFile{Mode: perm | os.ModeDir}
	s.FS[path] = &dir // MapFs needs a pointer
	return nil
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
