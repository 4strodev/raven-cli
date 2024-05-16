package template

import (
	"fmt"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/spf13/afero"
)

// TemplateConfig stores the config values available for a template
type TemplateConfig struct {
	SourceDirectory string `config:"source_directory"`
}

func loadTemplateConfigFromFile(fs afero.Fs, configFilePath string) (TemplateConfig, error) {
	stat, err := fs.Stat(configFilePath)
	if err != nil {
		return TemplateConfig{}, fmt.Errorf("error reading '%s'", configFilePath)
	}

	if stat.IsDir() {
		return TemplateConfig{}, fmt.Errorf("'%s' is not a file", configFilePath)
	}

	file, err := fs.Open(configFilePath)
	if err != nil {
		return TemplateConfig{}, fmt.Errorf("error opening '%s': %s", configFilePath, err)
	}

	content, err := afero.ReadAll(file)
	if err != nil {
		return TemplateConfig{}, fmt.Errorf("error reading '%s': %s", configFilePath, err)
	}

	if err := k.Load(rawbytes.Provider(content), json.Parser()); err != nil {
		return TemplateConfig{}, fmt.Errorf("error loading '%s'", configFilePath)
	}

	config := TemplateConfig{}
	err = k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "config"})
	if err != nil {
		return TemplateConfig{}, fmt.Errorf("error loading '%s'", configFilePath)
	}

	return config, nil
}
