package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPost(t *testing.T) {

	server := NewServer()

	// t.Run("correct file is returned based on path", func(t *testing.T) {
	// 	request := newNoteRequest("/note1")
	// 	response := httptest.NewRecorder()
	// 	server.ServeHTTP(response, request)
	// 	assertStatus(t, response.Code, http.StatusOK)
	// 	assertResponseBody(t, response.Body.String(), "web/content/note1.md")
	// })

	// t.Run("landing page returns a note", func(t *testing.T) {
	// 	request := newNoteRequest("/")
	// 	response := httptest.NewRecorder()
	// 	server.ServeHTTP(response, request)
	// 	assertStatus(t, response.Code, http.StatusOK)
	// 	assertResponseBody(t, response.Body.String(), "hello note server")
	// })

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
