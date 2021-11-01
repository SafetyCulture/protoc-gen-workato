package template

import (
	"github.com/SafetyCulture/protoc-gen-workato/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// WorkatoTemplate is an interface to use when rendering a workato connector
// https://docs.workato.com/developing-connectors/sdk/sdk-reference.html
type WorkatoTemplate struct {
	config *config.Config

	// All of the messages from the proto files
	messageMap map[string]*gendoc.Message
	// All of the enums from the proto files
	enumMap map[string]*gendoc.Enum

	// A map of the used messages from included methods
	usedMessageMap map[string]*gendoc.Message
	// A map of the used enums from the used messages
	usedEnumMap map[string]*gendoc.Enum
	// An ordered slice of the used messages from included methods
	messages []*gendoc.Message
	// An ordered slice of the used enums from the used messages
	enums []*gendoc.Enum

	// All of the included actions
	actions []*Action
	// A map of the actions grouped by their resource
	groupedActionMap map[string]*ActionGroup
	// An ordered slice of the grouped actions
	groupedActions []*ActionGroup

	// ObjectDefinitions are Workato formatted definitions of messages
	ObjectDefinitions []*ObjectDefinition
	// Actions are Workato formatted defintions of grouped methods
	Actions []*ActionDefinition
	// Picklists are Workato formatted definitions of enums and action groups
	Picklists []*PicklistDefinition
}

// FromGenDoc converts a protoc-gen-doc template to our template file
func FromGenDoc(template *gendoc.Template, cfg *config.Config) *WorkatoTemplate {
	workatoTemplate := &WorkatoTemplate{
		config:           cfg,
		messageMap:       make(map[string]*gendoc.Message),
		enumMap:          make(map[string]*gendoc.Enum),
		usedMessageMap:   make(map[string]*gendoc.Message),
		usedEnumMap:      make(map[string]*gendoc.Enum),
		groupedActionMap: make(map[string]*ActionGroup),
	}

	for _, file := range template.Files {
		// Create a index of all the used messages and enums so it's easy to reference later
		for _, message := range file.Messages {
			workatoTemplate.messageMap[message.FullName] = message
		}
		for _, enum := range file.Enums {
			workatoTemplate.enumMap[enum.FullName] = enum
		}

		// Find all of the actions we want to expose
		for _, service := range file.Services {
			for _, method := range service.Methods {
				isPublic := false
				if opts, ok := method.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation").(*options.Operation); ok {
					for _, tag := range opts.Tags {
						if tag == "Public" {
							isPublic = true
						}
					}
				}

				if isPublic {
					workatoTemplate.actions = append(workatoTemplate.actions, &Action{service, method})
				}
			}
		}
	}

	workatoTemplate.groupActions()
	workatoTemplate.generateObjectDefintions()
	workatoTemplate.generateActionDefinitions()
	workatoTemplate.generateEnumPicklists()

	return workatoTemplate
}
