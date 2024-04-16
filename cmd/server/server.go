package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	staticHandler http.Handler
}

func NewServer() *Server {
	fs := http.FileServer(http.Dir("./web/static"))
	staticHandler := http.StripPrefix("/static/", fs)
	return &Server{staticHandler: staticHandler}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "OOPs, we can't seem to find the page you're looking for")
		return
	}
	fmt.Fprint(w, "hello note server")
}
