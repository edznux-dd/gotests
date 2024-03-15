package templates

import (
	"embed"
	"io/fs"
)

//go:embed test/*.tmpl
var TestFS embed.FS

//go:embed testify/*.tmpl
var TestifyFS embed.FS

//go:embed fuzz/*.tmpl
var FuzzFS embed.FS

var TemplatesToFS = map[string]fs.FS{
	"fuzz":    FuzzFS,
	"testify": TestifyFS,
	"test":    TestFS,
}
