package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/nkawaller/notes/internal/utils"
	"github.com/nkawaller/notes/testutils"
)

func TestSSG(t *testing.T) {

	mockFS := testutils.StubFS
	ssg := NewStaticSiteGenerator(mockFS)

	t.Run("createStaticDir() creates the deploy directory", func(t *testing.T) {
		assertNotExist(t, mockFS, "deploy/static")
		ssg.createStaticDir()
		assertExists(t, mockFS, "deploy/static")
	})

	t.Run("HTML files are created from existing markdown files", func(t *testing.T) {
		assertExists(t, mockFS, "web/content/index.md")
		assertNotExist(t, mockFS, "index.html")
		ssg.generateStaticSite()
		assertExists(t, mockFS, "deploy/index.html")
	})

	t.Run("Root page (note index) is created in correct location", func(t *testing.T) {
		assertNotExist(t, mockFS, "web/content/root.md")
		assertNotExist(t, mockFS, "deploy/root.html")
		ssg.generateRootPage()
		assertExists(t, mockFS, "web/content/root.md")
		ssg.generateStaticSite()
		assertExists(t, mockFS, "deploy/root.html")
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
