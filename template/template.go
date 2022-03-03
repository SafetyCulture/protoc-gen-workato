package template

import (
	"fmt"
	"strings"

	"github.com/SafetyCulture/protoc-gen-workato/config"
	workato "github.com/SafetyCulture/protoc-gen-workato/s12/protobuf/workato"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"google.golang.org/genproto/googleapis/api/visibility"
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

	// A map of visibility rules that should be included in the connector
	visibilityRulesMap map[string]bool

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
	// AppBaseURl is the base URL to connect to for the application
	AppBaseURL string
	// DeveloperDocsURL is the URL to view the API docs on
	DeveloperDocsURL string

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
		visibilityRulesMap: make(map[string]bool),
		messageMap:         make(map[string]*gendoc.Message),
		enumMap:            make(map[string]*gendoc.Enum),
		usedMessageMap:     make(map[string]*gendoc.Message),
		usedEnumMap:        make(map[string]*gendoc.Enum),
		groupedActionMap:   make(map[string]*ActionGroup),
		dynamicPicklistMap: make(map[string]*schema.PicklistDefinition),
		Name:               cfg.Name,
		AppBaseURL:         cfg.AppBaseURL,
		DeveloperDocsURL:   cfg.DeveloperDocsURL,
		Methods:            cfg.CustomMethods,
	}

	for _, rule := range cfg.VisibilityRules {
		workatoTemplate.visibilityRulesMap[rule] = true
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

				if workatoTemplate.checkVisibility(method.Option("google.api.method_visibility")) {
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

func (t *WorkatoTemplate) checkVisibility(r interface{}) bool {
	isVisible := true

	if rule, ok := r.(*visibility.VisibilityRule); ok && rule != nil {
		restrictions := strings.Split(strings.TrimSpace(rule.Restriction), ",")

		if len(restrictions) != 0 {
			isVisible = false
		}
		for _, restriction := range restrictions {
			if t.visibilityRulesMap[strings.TrimSpace(restriction)] {
				isVisible = true
				break
			}
		}
	}

	return isVisible
}
