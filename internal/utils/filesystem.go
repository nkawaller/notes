package utils

import (
	"io"
	"os"
)

type FileSystem interface {
	Create(name string) (io.Writer, error)
	MkdirAll(path string, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
	ReadDir(name string) ([]os.DirEntry, error)
	Stat(name string) (os.FileInfo, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
	GetContentRoot() string
	GetTemplateLocation() string
}

type DefaultFileSystem struct {
	ContentRoot      string
	TemplateLocation string
}

func (fs DefaultFileSystem) Create(name string) (io.Writer, error) {
	return os.Create(name)
}

func (fs DefaultFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs DefaultFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (fs DefaultFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func (fs DefaultFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (fs DefaultFileSystem) WriteFile(name string, data []byte, perm os.FileMode) error {
	err := os.WriteFile(name, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func (fs DefaultFileSystem) GetContentRoot() string {
	return fs.ContentRoot
}

func (fs DefaultFileSystem) GetTemplateLocation() string {
	return fs.TemplateLocation
}
