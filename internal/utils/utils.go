package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Depado/bfchroma/v2"
	"github.com/russross/blackfriday/v2"
)

func GetMarkdownFilePath(path string) string {
	if path == "" {
		return filepath.Join("..", "..", "web", "content", "index.md")
	}
	return filepath.Join("..", "..", "web", "content", fmt.Sprintf("%s.md", path))
}

func ReadMarkdownFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func ConvertMarkdownToHTML(content []byte) []byte {
	var exts = blackfriday.NoIntraEmphasis | blackfriday.Tables | blackfriday.FencedCode | blackfriday.Autolink |
		blackfriday.Strikethrough | blackfriday.SpaceHeadings | blackfriday.BackslashLineBreak |
		blackfriday.DefinitionLists | blackfriday.Footnotes

	return blackfriday.Run([]byte(content), blackfriday.WithRenderer(bfchroma.NewRenderer(bfchroma.Style("gruvbox"))),
		blackfriday.WithExtensions(exts),
	)
}
