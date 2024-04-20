package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

type StubFileSystem struct {
	fs fstest.MapFS
	ContentRoot string
}

func (s StubFileSystem) ReadFile(filename string) ([]byte, error) {
	return s.fs.ReadFile(filename)
}

func (s StubFileSystem) ContentRootFn() string {
	return s.ContentRoot
}

func TestGETPost(t *testing.T) {


	fs := fstest.MapFS{
		"index.md": {Data: []byte("hello INDEX")},
		"test.md": {Data: []byte("hello TEST")},
	}

	stubFileSystem := StubFileSystem{fs: fs, ContentRoot: "/web/content/"}
	server := NewServer(stubFileSystem)

	t.Run("landing page returns a note", func(t *testing.T) {
		request := newNoteRequest("/test")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "hello note server")
	})

// 	t.Run("index route returns a note", func(t *testing.T) {
// 		request := newNoteRequest("/")
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertResponseByteLength(t, response.Body.Len(), 389)
// 	})

// 	t.Run("hello route produces ascii values for html-formatted H1 hello", func(t *testing.T) {
// 		request := newNoteRequest("/hello")
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assertStatus(t, response.Code, http.StatusOK)
// 		fmt.Println(response.Body.String())
// 		want := "[60 104 49 62 72 101 108 108 111 60 47 104 49 62 10]"
// 		assertSameBytes(t, response.Body.String(), want)
// 	})

// 	t.Run("unknown route returns a 404", func(t *testing.T) {
// 		request := newNoteRequest("/unknown")
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assertStatus(t, response.Code, http.StatusNotFound)
// 		assertResponseBody(t, response.Body.String(), "404 page not found\n")
// 	})
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
