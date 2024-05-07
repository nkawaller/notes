package testutils

import (
	"fmt"
	"bytes"
	"testing/fstest"
	"io/fs"
	"os"
)

type StubFileSystem struct {
	FS               fstest.MapFS
	ContentRoot      string
	TemplateLocation string
}

func (s StubFileSystem) Create(name string) (*os.File, error) {
	// Create a buffer to store the content
	buf := new(bytes.Buffer)
	// Add the buffer as a file to the MapFS
	s.FS[name] = &fstest.MapFile{
		Data: buf.Bytes(),
	}
	// Return a new file object
	// file := &os.File{}
	//
	fmt.Printf("Name is >>>>>>>> %s\n", name)
	fmt.Println("Filesystem name >>>>>>> ", s.FS[name])
	fmt.Println("Filesystem >>>>>>> ", s.FS)


	return nil, nil
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
