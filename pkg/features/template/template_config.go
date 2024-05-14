package template

// TemplateConfig stores the config values available for a template
type TemplateConfig struct {
	SourceDirectory string `koanf:"source_directory"`
}
