package main

import (
	"os"
	"testing"

	"github.com/nkawaller/notes/internal/utils"
	"github.com/nkawaller/notes/testutils"
)

func TestSSG(t *testing.T) {

	mockFS := testutils.StubFS
	ssg := NewStaticSiteGenerator(mockFS)

	t.Run("Static site generator creates the deploy directory", func(t *testing.T) {
		assertNotExist(t, mockFS, "deploy/static")
		ssg.generateStaticSite()
		assertExists(t, mockFS, "deploy/static")
	})

	t.Run("HTML files are created from existing markdown files", func(t *testing.T) {
		assertExists(t, mockFS, "web/content/index.md")
		assertNotExist(t, mockFS, "index.html")
		ssg.generateStaticSite()
		assertExists(t, mockFS, "deploy/index.html")
	})
}

func assertNotExist(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); !os.IsNotExist(err) {
		t.Errorf("%s should not exist", path)
	}
}

func assertExists(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); err != nil {
		t.Errorf("%s should exist", path)
	}
}
