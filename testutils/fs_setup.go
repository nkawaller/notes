package testutils

import (
    "os"
    "testing/fstest"
    "time"
)

var (

    indexModTime, _ = time.Parse(time.RFC3339, "2023-10-30T12:00:00Z")
	baconModTime, _ = time.Parse(time.RFC3339, "2024-11-11T12:00:00Z")
	mockTemplate, _ = os.ReadFile("../../testdata/mock_template.html")

	fs = fstest.MapFS{
		"index.md":           {Data: []byte("INDEX PAGE"), ModTime: indexModTime},
		"bacon.md":           {Data: []byte("BACON"), ModTime: baconModTime},
		"base_template.html": {Data: []byte(mockTemplate)},
	}

	StubFS = StubFileSystem{
		FS:               fs,
		ContentRoot:      "",
		TemplateLocation: "base_template.html",
    }
)
