package main

import (
	"log"
	"net/http"
)

func main() {
	server := &Server{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
