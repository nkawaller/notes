package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
