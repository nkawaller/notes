package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Depado/bfchroma/v2"
	"github.com/nkawaller/notes/internal/page"
	"github.com/russross/blackfriday/v2"
)

func GetMarkdownFilePath(fs FileSystem, endpoint string) string {
	if endpoint == "" {
		return filepath.Join(fs.ContentRootFn(), "index.md")
	}
	return filepath.Join(fs.ContentRootFn(), fmt.Sprintf("%s.md", endpoint))
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

func RenderPage(w http.ResponseWriter, html []byte, fs FileSystem, markdownFile string) {
	page := CreatePage(html, fs, markdownFile)
	ExecuteTemplate(w, page)
}

func CreatePage(html []byte, fs FileSystem, markdownFile string) page.Page {

	// Get the file's last modified date
	fileInfo, err := fs.Stat(markdownFile)
	if err != nil {
		log.Fatal(err)
	}
	lastModified := fileInfo.ModTime()

	// Create a Page struct
	page := page.Page{
		Title: "Note",
		HTML:  template.HTML(html),
		LastModified: lastModified,
		CSSPath: "../static/output.css",
	}
	return page
}

func ExecuteTemplate(w http.ResponseWriter, page page.Page) {

	// Read the base template file
	baseTemplate, err := os.ReadFile("../../web/templates/base_template.html")
	if err != nil {
		log.Fatal(err)
	}

	// Parse and execute the tempalte
	tmpl, err := template.New("base").Parse(string(baseTemplate))
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal(err)
	}
}
