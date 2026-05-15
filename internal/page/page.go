package page

import (
	"html/template"
	"time"
)

type Backlink struct {
	Slug  string
	Title string
}

type Page struct {
	Title        string
	HTML         template.HTML
	LastModified time.Time
	CSSPath      string
	ICONPath     string
	Backlinks    []Backlink
}
