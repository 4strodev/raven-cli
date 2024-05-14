package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/4strodev/raven/pkg/features/template"
	"github.com/spf13/afero"
)

func CreateProject(fs afero.Fs, targetDirectory string, templateName string) error {
	var userHomePath, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot get user home directory: %s", err)
	}

	var templatePath = filepath.Clean(filepath.Join(userHomePath, templateName))

	template, err := template.LoadFromFs(fs, templatePath)
	if err != nil {
		return fmt.Errorf("error loading template: %s", err)
	}


	
	return nil
}
