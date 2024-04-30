package main

import (
	"fmt"
	"log"
)

func main() {
	ssg := NewStaticSiteGenerator()
	err := ssg.generateStaticSite()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Static site generated successfully.")
}
