package config

import "github.com/SafetyCulture/protoc-gen-workato/template/schema"

// ConfigMethod is the overwritten code for a specific method
type ConfigMethod struct {
	Exec string `yaml:"exec"`
}

type ConfigMessage struct {
	Exec string `yaml:"exec"`
}

type ConfigAction struct {
	InputFields []schema.FieldDefinition `yaml:"input_fields"`
}

// Config is the configuration of the plugin
type Config struct {
	TemplateFile  string
	Action        map[string]ConfigAction    `yaml:"action"`
	Method        map[string]ConfigMethod    `yaml:"method"`
	Message       map[string]ConfigMessage   `yaml:"message"`
	CustomMethods []*schema.MethodDefinition `yaml:"custom_methods"`
}
