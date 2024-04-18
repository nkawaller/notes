package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nkawaller/notes/internal/utils"
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

	s.router.Handle("/", http.HandlerFunc(s.handleRequest))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	filename := strings.TrimPrefix(path, "/")
	markdownFile := utils.GetMarkdownFilePath(filename)
	content, err := utils.ReadMarkdownFile(markdownFile)

	if os.IsNotExist(err) {
		http.NotFound(w, r)
		// TODO: render a 404 template
		return
	} else if err != nil {
		log.Fatal(err)
	}

	html := utils.ConvertMarkdownToHTML(content)
	fmt.Println(html)
}
