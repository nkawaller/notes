package utils

import (
	"os"
)

type FileSystem interface {
	Create(name string) (*os.File, error)
	MkdirAll(path string, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
    ReadDir(name string) ([]os.DirEntry, error)
	Stat(name string) (os.FileInfo, error)
	GetContentRoot() string
	GetTemplateLocation() string
}

type DefaultFileSystem struct {
	ContentRoot      string
	TemplateLocation string
}

func (fs DefaultFileSystem) Create(name string) (*os.File, error) {
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

func (fs DefaultFileSystem) GetContentRoot() string {
	return fs.ContentRoot
}

func (fs DefaultFileSystem) GetTemplateLocation() string {
	return fs.TemplateLocation
}
