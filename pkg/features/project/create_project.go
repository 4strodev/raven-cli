package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/4strodev/raven/pkg/features/template"
	"github.com/spf13/afero"
)

// Creates a project given a template name and a target directory
func CreateProject(fs afero.Fs, targetDirectory string, templateName string) error {
	var userHomePath, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot get user home directory: %s", err)
	}

	var templatePath = filepath.Clean(filepath.Join(userHomePath, ".raven", "templates", templateName))

	template, err := template.LoadFromFs(fs, templatePath)
	if err != nil {
		return fmt.Errorf("error loading template: %s", err)
	}

	afero.DirExists(fs, targetDirectory)

	fmt.Printf("Loaded template %v\n", template)
	
	return nil
}
