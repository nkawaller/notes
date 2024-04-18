package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetMarkdownFilePath(path string) string {
	if path == "" {
		return filepath.Join("web", "content", "index.md")
	}
	return filepath.Join("web", "content", fmt.Sprintf("%s.md", path))
}

func ReadMarkdownFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
