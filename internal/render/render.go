package render

import (
	"fmt"
	"io"
	"io/fs"
	"text/template"

	"github.com/cweill/gotests/internal/models"
)

type Render struct {
	tmpls *template.Template
}

func New() *Render {
	r := Render{
		tmpls: template.New("render").Funcs(map[string]interface{}{
			"Field":                fieldName,
			"Receiver":             receiverName,
			"Param":                parameterName,
			"Want":                 wantName,
			"Got":                  gotName,
			"DefaultValueFromType": defaultValueForType,
		}),
	}

	var err error
	r.tmpls, err = r.tmpls.ParseFS(DefaultTemplates, "*/*.tmpl")
	if err != nil {
		// if there's an error here, we have a problem in the default template
		// This should basically be a compile time error, so panic'ing to error asap.
		panic(err)
	}

	return &r
}

// LoadCustomTemplates allows to load in custom templates from a specified path.
func (r *Render) LoadCustomTemplates(fs fs.FS) error {
	var err error
	r.tmpls, err = r.tmpls.ParseFS(fs, "*/*.tmpl")
	if err != nil {
		return fmt.Errorf("LoadCustomTemplates: %w", err)
	}

	return nil
}

// LoadFromData allows to load from a data slice
func (r *Render) LoadFromData(templateData [][]byte) {
	for _, d := range templateData {
		r.tmpls = template.Must(r.tmpls.Parse(string(d)))
	}
}

func (r *Render) Header(w io.Writer, h *models.Header) error {
	if err := r.tmpls.ExecuteTemplate(w, "header", h); err != nil {
		return err
	}
	_, err := w.Write(h.Code)
	return err
}

func (r *Render) TestFunction(
	w io.Writer,
	f *models.Function,
	printInputs bool,
	subtests bool,
	named bool,
	parallel bool,
	params map[string]interface{}) error {
	return r.tmpls.ExecuteTemplate(w, "function", struct {
		*models.Function
		PrintInputs    bool
		Subtests       bool
		Parallel       bool
		Named          bool
		TemplateParams map[string]interface{}
	}{
		Function:       f,
		PrintInputs:    printInputs,
		Subtests:       subtests,
		Parallel:       parallel,
		Named:          named,
		TemplateParams: params,
	})
}
