package main

import (
	"fmt"
	"log"

	"github.com/nkawaller/notes/internal/utils"
)

func main() {
	contentRoot := "web/content/"
	templateLocation := "web/templates/base_template.html"
	fileSystem := utils.DefaultFileSystem{ContentRoot: contentRoot, TemplateLocation: templateLocation}
	ssg := NewStaticSiteGenerator(fileSystem)
	ssg.generateRootPage()
	err := ssg.generateStaticSite()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Static site generated successfully.")
}
