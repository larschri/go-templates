package web

import (
	"embed"
	"html/template"
)

var (
	//go:embed static
	Static embed.FS

	//go:embed templates/*.html
	TemplatesFS embed.FS

	Templates = template.Must(template.ParseFS(TemplatesFS, "templates/*"))
)
