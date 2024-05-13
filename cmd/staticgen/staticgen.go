package main

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
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

func (s *StaticSiteGenerator) generateRootPage() {

	files, err := s.fileSystem.ReadDir(s.fileSystem.GetContentRoot())
	if err != nil {
		log.Fatal(err)
	}

	var markdownFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") && file.Name() != "rood.md" {
			markdownFiles = append(markdownFiles, file.Name())
		}
	}

	markdownContent := utils.GenerateMarkdownContent(markdownFiles, s.fileSystem)
	outputFile := filepath.Join(s.fileSystem.GetContentRoot(), "root.md")

	err = s.fileSystem.WriteFile(outputFile, []byte(markdownContent), 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Markdown file generated and saved to %s\n", outputFile)

}

func (s *StaticSiteGenerator) generateStaticSite() error {

	if err := s.createStaticDir(); err != nil {
		return err
	}

	files, err := s.fileSystem.ReadDir(s.fileSystem.GetContentRoot())
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
		err := s.processMarkdownFile(file, tmpl)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (s *StaticSiteGenerator) createStaticDir() error {
	err := s.fileSystem.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *StaticSiteGenerator) processMarkdownFile(file fs.DirEntry, tmpl *template.Template) error {
	if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
		path := file.Name()
		content, err := utils.ReadMarkdownFile(s.fileSystem, filepath.Join(s.fileSystem.GetContentRoot(), path))
		if err != nil {
			return err
		}

		html := utils.ConvertMarkdownToHTML(content)

		fileInfo, err := s.fileSystem.Stat(filepath.Join(s.fileSystem.GetContentRoot(), path))
		if err != nil {
			return err
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
			return err
		}

		defer func() {
			if closer, ok := outputFile.(io.WriteCloser); ok {
				closer.Close()
			}
		}()

		err = tmpl.Execute(outputFile, page)
		if err != nil {
			return err
		}
	}
	return nil
}
