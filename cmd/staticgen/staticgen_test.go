package main

import (
	"os"
	"testing"

	"github.com/nkawaller/notes/internal/utils"
	"github.com/nkawaller/notes/testutils"
)

func TestSSG(t *testing.T) {

	fs := testutils.StubFS
	ssg := NewStaticSiteGenerator(fs)
	dirPath := "deploy/static"

	t.Run("Static site generator creates the deploy directory", func(t *testing.T) {
		assertDirNotExist(t, fs, dirPath)
		ssg.generateStaticSite()
		assertDirExists(t, fs, dirPath)

	})
}

func assertDirNotExist(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); !os.IsNotExist(err) {
		t.Errorf("Directory %s should not exist initially", path)
	}
}

func assertDirExists(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); err != nil {
		t.Errorf("Directory %s should exist after generating static site", path)
	}
}
