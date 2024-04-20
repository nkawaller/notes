package utils

import "os"

type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
}

type DefaultFileSystem struct{}

func (fs DefaultFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
