package gowiki

import "html/template"

// Page represents the fields composing a wiki page.
type Page struct {
	Title string
	Body  string
}

// PageService is the interface for manipulating Pages
type PageService interface {
	LoadPage(title string) (*Page, error)
	Save(p *Page) error
}

type Renderer interface {
	ToHtml(content string) template.HTML
}
