package utils

import (
	"os"
)

type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	ContentRootFn() string
}

type DefaultFileSystem struct {
	ContentRoot string
}

func (fs DefaultFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (fs DefaultFileSystem) ContentRootFn() string {
	return fs.ContentRoot
}
