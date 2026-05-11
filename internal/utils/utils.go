package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Depado/bfchroma/v2"
	"github.com/nkawaller/notes/internal/page"
	"github.com/russross/blackfriday/v2"
)

func GetMarkdownFilePath(fs FileSystem, endpoint string) string {
	if endpoint == "" {
		return filepath.Join(fs.GetContentRoot(), "index.md")
	}
	return filepath.Join(fs.GetContentRoot(), fmt.Sprintf("%s.md", endpoint))
}

func ReadMarkdownFile(fs FileSystem, filePath string) ([]byte, error) {
	return fs.ReadFile(filePath)
}

func ConvertMarkdownToHTML(content []byte) []byte {
	var exts = blackfriday.NoIntraEmphasis | blackfriday.Tables | blackfriday.FencedCode | blackfriday.Autolink |
		blackfriday.Strikethrough | blackfriday.SpaceHeadings | blackfriday.BackslashLineBreak |
		blackfriday.DefinitionLists | blackfriday.Footnotes

	return blackfriday.Run([]byte(content), blackfriday.WithRenderer(bfchroma.NewRenderer(bfchroma.Style("gruvbox"))),
		blackfriday.WithExtensions(exts),
	)
}

func RenderPage(w http.ResponseWriter, html []byte, fs FileSystem, markdownFile string) error {
	p, err := CreatePage(html, fs, markdownFile)
	if err != nil {
		return err
	}
	return ExecuteTemplate(w, fs, p)
}

func CreatePage(html []byte, fs FileSystem, markdownFile string) (page.Page, error) {
	fileInfo, err := fs.Stat(markdownFile)
	if err != nil {
		return page.Page{}, err
	}
	p := page.Page{
		Title:        "Note",
		HTML:         template.HTML(html),
		LastModified: fileInfo.ModTime(),
		CSSPath:      "../static/output.css",
	}
	return p, nil
}

func ExecuteTemplate(w http.ResponseWriter, fs FileSystem, page page.Page) error {
	baseTemplate, err := fs.ReadFile(fs.GetTemplateLocation())
	if err != nil {
		return err
	}
	tmpl, err := template.New("base").Parse(string(baseTemplate))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, page)
}

func GenerateMarkdownContent(files []string, fs FileSystem) (string, error) {
	sort.Strings(files)

	var content strings.Builder

	content.WriteString("# All notes\n\n")
	content.WriteString("|      |                     |\n")
	content.WriteString("|------|---------------------|\n")

	for _, file := range files {
		fileInfo, err := fs.Stat(filepath.Join(fs.GetContentRoot(), file))
		if err != nil {
			return "", err
		}
		modTime := fileInfo.ModTime().Year()
		link := fmt.Sprintf("| %d | [%s][] |\n", modTime, strings.TrimSuffix(file, ".md"))
		content.WriteString(link)
	}

	content.WriteString("\n")

	for _, file := range files {
		link := fmt.Sprintf("[%s]: %s\n", strings.TrimSuffix(file, ".md"), strings.TrimSuffix(file, ".md")+".html")
		content.WriteString(link)
	}

	return content.String(), nil
}
