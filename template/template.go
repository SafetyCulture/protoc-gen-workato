package template

import (
	"github.com/SafetyCulture/protoc-gen-workato/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

type WorkatoTemplate struct {
	config *config.Config

	messageMap map[string]*gendoc.Message
	enumMap    map[string]*gendoc.Enum

	usedMessageMap map[string]*gendoc.Message
	usedEnumMap    map[string]*gendoc.Enum

	actions          []*Action
	groupedActionMap map[string]*ActionGroup
	groupedActions   []*ActionGroup

	Messages          []*gendoc.Message
	Enums             []*gendoc.Enum
	ObjectDefinitions []*ObjectDefinition
	Actions           []*ActionDefinition
	Picklists         []*PicklistDefinition
}

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

	return workatoTemplate
}
