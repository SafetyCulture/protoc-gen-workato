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
}

// Config is the configuration of the plugin
type Config struct {
	// Name is the name of the connector
	Name string `yaml:"name"`
	// AppBaseURL is the base URL to use when making API requests
	AppBaseURL string `yaml:"app_base_url"`
	// DeveloperDocsURL is the link to the developer documentation
	DeveloperDocsURL string `yaml:"developer_docs_url"`
	// TemplateFile is a custom template file for generating the connector
	TemplateFile string `yaml:"template_file"`
	// Action is a map of custom actions
	Action map[string]Action `yaml:"action"`
	// Method is a map of method overrides
	Method map[string]Method `yaml:"method"`
	// Message is a map of message overrides
	Message map[string]Message `yaml:"message"`
	// CustomMethods is a set of custom methods to push into the connector
	CustomMethods []*schema.MethodDefinition `yaml:"custom_methods"`
	// VisibilityRules define the set of google.api.VisibilityRule that are included in connector generation
	// By default APIs, Fields, Enum Values with no VisibilityRule will always be included, adding a restriction like
	// INTERNAL will prevent it from being generated, including INTERNAL in this slice will begin generating it again.
	VisibilityRules []string
}
