package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nkawaller/notes/internal/utils"
)

type Server struct {
	staticHandler http.Handler
	router        *http.ServeMux
	fileSystem    utils.FileSystem
}

func NewServer(fileSystem utils.FileSystem) *Server {
	staticHandler := http.FileServer(http.Dir("web/static"))

	s := &Server{
		staticHandler: staticHandler,
		router:        http.NewServeMux(),
		fileSystem:    fileSystem,
	}

	s.router.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	s.router.Handle("/", http.HandlerFunc(s.handleRequest))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	filename := strings.TrimPrefix(path, "/")
	markdownFile := utils.GetMarkdownFilePath(s.fileSystem, filename)
	content, err := utils.ReadMarkdownFile(s.fileSystem, markdownFile)

	if os.IsNotExist(err) {
		http.NotFound(w, r)
		// TODO: render a 404 template
		return
	} else if err != nil {
		log.Fatal(err)
	}

	html := utils.ConvertMarkdownToHTML(content)
	utils.RenderPage(w, html, s.fileSystem, markdownFile)
}
