package main

import (
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/nkawaller/notes/testutils"
	"github.com/nkawaller/notes/internal/utils"
)

func TestSSG(t *testing.T) {

	indexModTime, _ := time.Parse(time.RFC3339, "2023-10-30T12:00:00Z")
	baconModTime, _ := time.Parse(time.RFC3339, "2024-11-11T12:00:00Z")
	mockTemplate, _ := os.ReadFile("../../testdata/mock_template.html")

	fs := fstest.MapFS{
		"index.md":           {Data: []byte("INDEX PAGE"), ModTime: indexModTime},
		"bacon.md":           {Data: []byte("BACON"), ModTime: baconModTime},
		"base_template.html": {Data: []byte(mockTemplate)},
	}

	stubFileSystem := testutils.StubFileSystem{
		FS:               fs,
		ContentRoot:      "",
		TemplateLocation: "base_template.html",
	}

	ssg := NewStaticSiteGenerator(stubFileSystem)
	dirPath := "deploy/static"

	t.Run("Static site generator creates the deploy directory", func(t *testing.T) {
		assertDirNotExist(t, stubFileSystem, dirPath)
		ssg.generateStaticSite()
		assertDirExists(t, stubFileSystem, dirPath)

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
