package config

type ConfigAction struct {
	// ID   string
	// Name string
	Exec string `yaml:"exec"`
}

// The configuration of the plugin
type Config struct {
	Name   string                  `yaml:"name"`
	Method map[string]ConfigAction `yaml:"method"`
}
