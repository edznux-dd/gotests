package render

import (
	"embed"
)

//go:embed templates/*.tmpl
var DefaultTemplates embed.FS
