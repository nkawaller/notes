package main

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
    "strings"

	"github.com/nkawaller/notes/internal/page"
	"github.com/nkawaller/notes/internal/utils"
)

type StaticSiteGenerator struct {
	fileSystem utils.FileSystem
}

func NewStaticSiteGenerator(fileSystem utils.FileSystem) *StaticSiteGenerator {
	s := &StaticSiteGenerator{
		fileSystem: fileSystem,
	}
	return s
}

func (s *StaticSiteGenerator) generateStaticSite() error {

	err := s.fileSystem.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}

	files, err := s.fileSystem.ReadDir(s.fileSystem.GetContentRoot())
	fmt.Println(s.fileSystem)
	if err != nil {
		log.Fatal(err)
	}

	baseTemplate, err := s.fileSystem.ReadFile(s.fileSystem.GetTemplateLocation())
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("base").Parse(string(baseTemplate))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			path := file.Name()
			content, err := utils.ReadMarkdownFile(s.fileSystem, filepath.Join(s.fileSystem.GetContentRoot(), path))
			if err != nil {
				log.Fatal(err)
			}

			html := utils.ConvertMarkdownToHTML(content)

			fileInfo, err := s.fileSystem.Stat(filepath.Join(s.fileSystem.GetContentRoot(), path))
			if err != nil {
				log.Fatal(err)
			}
			lastModified := fileInfo.ModTime()

			page := page.Page{
				Title:        "Markdown Blog",
				HTML:         template.HTML(html),
				LastModified: lastModified,
				CSSPath:      "./static/output.css",
			}

			outputPath := filepath.Join("deploy", strings.TrimSuffix(path, ".md")+".html")
			outputFile, err := s.fileSystem.Create(outputPath)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			err = tmpl.Execute(outputFile, page)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return nil
}
