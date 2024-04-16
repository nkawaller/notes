package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPost(t *testing.T) {

	server := &Server{}

	t.Run("landing page returns a note", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/"), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
