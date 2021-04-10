package operator

import (
	"io"
	"log"
	"text/template"
)

type Renderer interface {
	Render(w io.Writer, data interface{})
}

type renderer struct {
	tmpl *template.Template
}

func NewTemplateRenderer(pattern string) Renderer {
	tmpl := template.New("template")
	tmpl, err := tmpl.Option("missingkey=error").Parse(pattern)

	if err != nil {
		log.Fatalf("NewTemplateRenderer() could not create template: %v", err)
		return nil
	}

	return &renderer{tmpl}
}

func (r *renderer) Render(w io.Writer, data interface{}) {
	r.tmpl.Execute(w, data)
}
