package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer() *TemplateRenderer {
	templates, err := template.ParseGlob("cmd/templates/*.html")
	if err != nil {
		panic(err)
	}
	return &TemplateRenderer{
		Templates: templates,
	}
}
