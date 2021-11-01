package config

// ConfigAction is the overwritten code for a specific action
type ConfigAction struct {
	Exec string `yaml:"exec"`
}

// Config is the configuration of the plugin
type Config struct {
	TemplateFile string
	Method       map[string]ConfigAction `yaml:"method"`
}
