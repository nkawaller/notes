package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	staticHandler http.Handler
	router        *http.ServeMux
}

func NewServer() *Server {
	fs := http.FileServer(http.Dir("./web/static"))
	staticHandler := http.StripPrefix("/static/", fs)

	s := &Server{
		staticHandler,
		http.NewServeMux(),
	}

	s.router.Handle("/", http.HandlerFunc(s.indexHandler))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "OOPs, we can't seem to find the page you're looking for")
		return
	}
	fmt.Fprint(w, "hello note server")
}
