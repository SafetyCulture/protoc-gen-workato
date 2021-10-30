package config

type ConfigAction struct {
	// ID   string
	// Name string
	Exec string `yaml:"exec"`
}

// The configuration of the plugin
type Config struct {
	TemplateFile string
	Method       map[string]ConfigAction `yaml:"method"`
}
