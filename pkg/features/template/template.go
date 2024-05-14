package template

import (
	"errors"

	"github.com/spf13/afero"
)

type Template struct {
	config TemplateConfig
}

// Loads a template from a fs
func LoadFromFs(fs afero.Fs, templatePath string) (Template, error) {
	// TODO read template from fs
	return Template{}, errors.New("Not implemented")
}
