package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
    "os"
	"strings"
	"testing"
	"testing/fstest"
	"time"
)

type StubFileSystem struct {
	fs          fstest.MapFS
	ContentRoot string
    TemplateLocation string
}

func (s StubFileSystem) ReadFile(filename string) ([]byte, error) {
	return s.fs.ReadFile(filename)
}

func (s StubFileSystem) Stat(name string) (os.FileInfo, error) {
    return s.fs.Stat(name)
}

func (s StubFileSystem) ContentRootFn() string {
	return s.ContentRoot
}

func (s StubFileSystem) GetTemplateLocation() string {
	return s.TemplateLocation
}

func TestGETPost(t *testing.T) {

	indexModTime, _ := time.Parse(time.RFC3339, "2023-10-30T12:00:00Z")
	baconModTime, _ := time.Parse(time.RFC3339, "2024-11-11T12:00:00Z")
    mockTemplate, _ := os.ReadFile("../../testdata/mock_template.html")

	fs := fstest.MapFS{
        "index.md": {Data: []byte("INDEX PAGE"), ModTime: indexModTime},
        "bacon.md": {Data: []byte("BACON"), ModTime: baconModTime},
        "base_template.html": {Data: []byte(mockTemplate)},
	}

	// ContenRoot is an empty string here so we search for the file directly
    stubFileSystem := StubFileSystem{fs: fs, ContentRoot: "", TemplateLocation: "base_template.html"}
	server := NewServer(stubFileSystem)

	t.Run("Index page renders content correctly", func(t *testing.T) {
		request := newNoteRequest("/")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "<p>INDEX PAGE</p>")
	})

	t.Run("Index page renders last updated time correctly", func(t *testing.T) {
		request := newNoteRequest("/")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "Last updated 2023-10-30")
	})

	t.Run("Bacon page renders content correctly", func(t *testing.T) {
		request := newNoteRequest("/bacon")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "<p>BACON</p>")
	})

	t.Run("Bacon page renders content correctly", func(t *testing.T) {
		request := newNoteRequest("/bacon")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "Last updated 2024-11-11")
	})

	t.Run("Template contains correct header content", func(t *testing.T) {
		request := newNoteRequest("/")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "<a href=\"https://github.com/nkawaller\" class=\"ml-4\">Open Source</a>")
	})

	t.Run("Template contains correct footer content", func(t *testing.T) {
		request := newNoteRequest("/")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBodyContainsPattern(t, response.Body.String(), "<p><a href=\"https://magazine-b.com/\">한국어</a></p>")
	})

	t.Run("unknown route returns a 404", func(t *testing.T) {
		request := newNoteRequest("/unknown")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "404 page not found\n")
	})
}

func newNoteRequest(endpoint string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", endpoint), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body does not match; got %q, want %q", got, want)
	}
}

func assertResponseBodyContainsPattern(t *testing.T, body, pattern string) {
	if !strings.Contains(body, pattern) {
		t.Errorf("Couldn't find the following pattern: %s in response body", pattern)
	}
}

func assertResponseByteLength(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("Unexpected number of bytes; got %d, want %d", got, want)
	}
}

func assertSameBytes(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("Unexpected bytes; got %q, want %q", got, want)
	}
}
