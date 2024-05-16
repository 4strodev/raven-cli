package template

import (
	"fmt"
	"path/filepath"

	"github.com/knadh/koanf/v2"
	"github.com/spf13/afero"
)

const _TEMPLATE_MANIFEST = "raven.json"

var k = koanf.New(".")

type Template struct {
	config TemplateConfig
}

// Loads a template from a fs
func LoadFromFs(fs afero.Fs, templatePath string) (Template, error) {
	// Checking if template directory exists
	exists, err := afero.DirExists(fs, templatePath)
	if err != nil {
		return Template{}, fmt.Errorf("error checking directory existence: %s", templatePath)
	}
	if !exists {
		return Template{}, fmt.Errorf("directory '%s' does not exists", templatePath)
	}

	// Reading directory content
	files, err := afero.ReadDir(fs, templatePath)
	if err != nil {
		return Template{}, fmt.Errorf("error reading template directory '%s'", err)
	}

	// Find template manifest
	var found = false
	var templateConfigFilePath = ""
	for _, fileInfo := range files {
		if fileInfo.Name() != _TEMPLATE_MANIFEST {
			continue
		}
		templateConfigFilePath = filepath.Join(templatePath, fileInfo.Name())
		found = true
	}
	if !found {
		return Template{}, fmt.Errorf("'%s' not found", _TEMPLATE_MANIFEST)
	}

	// Read template manifest and extract config
	config, err := loadTemplateConfigFromFile(fs, templateConfigFilePath)
	if err != nil {
		return Template{}, err
	}

	return Template{
		config,
	}, nil
}
