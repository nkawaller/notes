package page

import (
	"html/template"
	"time"
)

type Page struct {
	Title        string
	HTML         template.HTML
	LastModified time.Time
	CSSPath      string
}
