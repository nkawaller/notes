package main

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"sort"
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
		if strings.HasSuffix(file.Name(), ".md") && file.Name() != "root.md" {
			markdownFiles = append(markdownFiles, file.Name())
		}
	}

	markdownContent, err := utils.GenerateMarkdownContent(markdownFiles, s.fileSystem)
	if err != nil {
		log.Fatal(err)
	}
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

	backlinks, err := s.buildBacklinkIndex(files)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := s.processMarkdownFile(file, tmpl, backlinks)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

var excludedSources = map[string]bool{
	"root.md":        true,
	"korean-root.md": true,
}

func (s *StaticSiteGenerator) buildBacklinkIndex(files []fs.DirEntry) (map[string][]page.Backlink, error) {
	noteSet := make(map[string]bool)
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".md" {
			noteSet[strings.TrimSuffix(f.Name(), ".md")] = true
		}
	}

	index := make(map[string][]page.Backlink)
	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".md" {
			continue
		}
		if excludedSources[f.Name()] {
			continue
		}

		sourceSlug := strings.TrimSuffix(f.Name(), ".md")
		content, err := s.fileSystem.ReadFile(filepath.Join(s.fileSystem.GetContentRoot(), f.Name()))
		if err != nil {
			return nil, err
		}

		title := utils.ExtractTitle(content)
		if title == "" {
			title = sourceSlug
		}

		seen := make(map[string]bool)
		for _, target := range utils.ExtractReferenceLinks(content) {
			if target == sourceSlug || !noteSet[target] || seen[target] {
				continue
			}
			seen[target] = true
			index[target] = append(index[target], page.Backlink{Slug: sourceSlug, Title: title})
		}
	}

	for target := range index {
		sort.Slice(index[target], func(i, j int) bool {
			return index[target][i].Title < index[target][j].Title
		})
	}
	return index, nil
}

func (s *StaticSiteGenerator) createStaticDir() error {
	err := s.fileSystem.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *StaticSiteGenerator) processMarkdownFile(file fs.DirEntry, tmpl *template.Template, backlinks map[string][]page.Backlink) error {
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

		slug := strings.TrimSuffix(path, ".md")
		page := page.Page{
			Title:        "Nathan Kawaller",
			HTML:         template.HTML(html),
			LastModified: lastModified,
			CSSPath:      "./static/output.css",
			ICONPath:     "./static/N.jpg",
			Backlinks:    backlinks[slug],
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
