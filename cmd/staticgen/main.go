package main

import (
	"fmt"
	"log"
)

func main() {
	err := generateStaticSite()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Static site generated successfully.")
}
