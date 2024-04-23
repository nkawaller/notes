package utils

import (
	"os"
)

type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	Stat(name string) (os.FileInfo, error)
	GetContentRoot() string
	GetTemplateLocation() string
}

type DefaultFileSystem struct {
	ContentRoot      string
	TemplateLocation string
}

func (fs DefaultFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
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
