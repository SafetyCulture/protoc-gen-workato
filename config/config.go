package config

import "github.com/SafetyCulture/protoc-gen-workato/template/schema"

// ConfigMethod allows for customization of a gRPC method
type ConfigMethod struct {
	Exec string `yaml:"exec"`
}

// ConfigMessage allows for customization of a gRPC message
type ConfigMessage struct {
	// Custom code for generating the object definition
	Exec string `yaml:"exec"`
}

// ConfigAction allows for customization of a grouped action
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
