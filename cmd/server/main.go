package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nkawaller/notes/internal/utils"
)

func main() {
	contentRoot := "../../web/content/"
	fileSystem := utils.DefaultFileSystem{ContentRoot: contentRoot}
	server := NewServer(fileSystem)

	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
