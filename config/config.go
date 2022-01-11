package config

import "github.com/SafetyCulture/protoc-gen-workato/template/schema"

// Method allows for customization of a gRPC method
type Method struct {
	Exec string `yaml:"exec"`
}

// Message allows for customization of a gRPC message
type Message struct {
	// Always include this message, even if it is not used by a method directly
	Include bool `yaml:"include"`
	// Custom code for generating the object definition
	Exec string `yaml:"exec"`
}

// Action allows for customization of a grouped action
type Action struct {
	InputFields        []schema.FieldDefinition `yaml:"input_fields"`
	DefaultHelpMessage *schema.HelpMessage      `yaml:"default_help_message"`
	Execute            *schema.Execute          `yaml:"execute"`
}

func (a *Action) GetExecute() *schema.Execute {
	if a.Execute == nil {
		return &schema.Execute{
			ExcludeKeys: []string{},
		}
	}
	return a.Execute
}

// Config is the configuration of the plugin
type Config struct {
	TemplateFile  string
	Action        map[string]Action          `yaml:"action"`
	Method        map[string]Method          `yaml:"method"`
	Message       map[string]Message         `yaml:"message"`
	CustomMethods []*schema.MethodDefinition `yaml:"custom_methods"`
}
