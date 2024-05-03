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

	t.Run("Static site generator creates the deploy directory", func(t *testing.T) {
		assertNotExist(t, fs, "deploy/static")
		ssg.generateStaticSite()
		assertExists(t, fs, "deploy/static")
	})

	// t.Run("HTML files are created from existing markdown files", func(t *testing.T) {
	// 	assertExists(t, fs, "index.md")
	// 	assertNotExist(t, fs, "index.html")
	// 	ssg.generateStaticSite()
 //        for f := range fs.FS {
 //            fmt.Println("File:", f)
 //        }
	// })
}

func assertNotExist(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); !os.IsNotExist(err) {
		t.Errorf("Directory %s should not exist initially", path)
	}
}

func assertExists(t testing.TB, fs utils.FileSystem, path string) {
	if _, err := fs.Stat(path); err != nil {
		t.Errorf("Directory %s should exist after generating static site", path)
	}
}
