package template

import (
	"fmt"

	"github.com/SafetyCulture/protoc-gen-workato/config"
	workato "github.com/SafetyCulture/protoc-gen-workato/s12/protobuf/workato"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// ServiceMethod is a combined service and method defintion
type ServiceMethod struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

// ExtractFirstTag Extract and Converts first non-public tag
func (t *ServiceMethod) extractFirstTag() (string, error) {
	opts, ok := t.Method.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation").(*options.Operation)
	if !ok {
		return "", fmt.Errorf("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation from method %s", t.Method.Name)
	}
	var tagName string
	for _, tag := range opts.Tags {
		if tag != "Public" {
			tagName = tag
			break
		}
	}

	if tagName == "" {
		return "", fmt.Errorf("couldn't find any tags for method %s", t.Method.Name)
	}

	return tagName, nil
}

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

	dynamicPicklistMap map[string]*schema.PicklistDefinition

	// All the included actions
	actions []*ServiceMethod
	// A map of the actions grouped by their resource
	groupedActionMap map[string]*ActionGroup
	// An ordered slice of the grouped actions
	groupedActions []*ActionGroup

	// Name is the name of the connector
	Name string

	// ObjectDefinitions are Workato formatted definitions of messages
	ObjectDefinitions []*schema.ObjectDefinition
	// Actions are Workato formatted defintions of grouped methods
	Actions []*schema.ActionDefinition
	// Picklists are Workato formatted definitions of enums and action groups
	Picklists []*schema.PicklistDefinition
	Methods   []*schema.MethodDefinition

	// All triggers
	triggers []*ServiceMethod
	// Triggers are Workato formatted definitions of grouped triggers
	Triggers []*schema.TriggerDefinition
}

// FromGenDoc converts a protoc-gen-doc template to our template file
func FromGenDoc(template *gendoc.Template, cfg *config.Config) (*WorkatoTemplate, error) {
	workatoTemplate := &WorkatoTemplate{
		config:             cfg,
		messageMap:         make(map[string]*gendoc.Message),
		enumMap:            make(map[string]*gendoc.Enum),
		usedMessageMap:     make(map[string]*gendoc.Message),
		usedEnumMap:        make(map[string]*gendoc.Enum),
		groupedActionMap:   make(map[string]*ActionGroup),
		dynamicPicklistMap: make(map[string]*schema.PicklistDefinition),
		Name:               cfg.Name,
		Methods:            cfg.CustomMethods,
	}

	for _, file := range template.Files {
		// Create an index of all the used messages and enums, so it's easy to reference later
		for _, message := range file.Messages {
			workatoTemplate.messageMap[message.FullName] = message
		}
		for _, enum := range file.Enums {
			removeUnspecifiedValue(enum)
			workatoTemplate.enumMap[enum.FullName] = enum
		}

		// Find all the actions we want to expose
		for _, service := range file.Services {
			for _, method := range service.Methods {
				seviceMethod := &ServiceMethod{service, method}

				workatoOpt, _ := method.Option("s12.protobuf.workato.method").(*workato.MethodOptionsWorkato)

				// Will exclude all the actions which have extension method with excluded flag set
				if workatoOpt != nil && workatoOpt.Excluded {
					continue
				}

				if workatoOpt != nil && workatoOpt.Picklist != nil {
					workatoTemplate.Picklists = append(workatoTemplate.Picklists, workatoTemplate.recordDynamicPicklist(seviceMethod, workatoOpt))
				}

				if workatoOpt != nil && workatoOpt.Trigger {
					workatoTemplate.triggers = append(workatoTemplate.triggers, seviceMethod)
					continue
				}

				isPublic := false
				if openapiOpt, ok := method.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation").(*options.Operation); ok {
					for _, tag := range openapiOpt.Tags {
						if tag == "Public" {
							isPublic = true
						}
					}
				}

				if isPublic {
					workatoTemplate.actions = append(workatoTemplate.actions, seviceMethod)
				}
			}
		}
	}

	workatoTemplate.groupActions()
	if err := workatoTemplate.generateTriggerDefinitions(); err != nil {
		return nil, err
	}
	workatoTemplate.captureIncludedMessages()
	workatoTemplate.generateObjectDefinitions()
	workatoTemplate.generateActionDefinitions()
	workatoTemplate.generateEnumPicklists()

	return workatoTemplate, nil
}
