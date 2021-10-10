package genworkato

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	tmpl "text/template"

	"github.com/Masterminds/sprig"
	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

func Escape(s string) string {
	return tmpl.HTMLEscapeString(s)
}

func EscapeActionName(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, ".", "_"), "/", "__")
}

func formatStringSlice(slc []string) string {
	if slc == nil {
		return "[]"
	}
	b, err := json.Marshal(slc)
	if err != nil {
		return "[]"
	}
	return string(b)
}

var funcMap = tmpl.FuncMap{
	"escape":            Escape,
	"escapeActionName":  EscapeActionName,
	"formatStringSlice": formatStringSlice,
}

type FieldDefinition struct {
	Name               string
	Label              string
	Optional           bool
	Type               string
	Hint               string
	Of                 string
	PropertiesRef      string
	Properties         []*ObjectDefinition
	ControlType        string
	ToggleHint         string
	ToggleField        *ObjectDefinition
	Default            string
	Picklist           string
	Delimiter          string
	Sticky             bool
	RenderInput        string
	ParseOutput        string
	ChangeOnBlur       bool
	SupportPills       bool
	Custom             bool
	ExtendsSchema      bool
	ListMode           string
	ListModeToggle     bool
	ItemLabel          string
	AddFieldLabel      string
	EmptySchemaMessage string
	SampleDataType     string
	NgIf               string
}

type ObjectDefinition struct {
	Name string

	Fields []*FieldDefinition
}

type Method struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
	Action  string
}

type Endpoint struct {
	ExcludeFromQuery []string
	Func             string
}

type ActionDefinition struct {
	Name        string
	Title       string
	Subtitle    string
	Description string

	ConfigFields []*FieldDefinition
	InputFields  map[string]string
	OutputFields map[string]string
	Endpoints    map[string]Endpoint
}

type PicklistValue struct {
	Key   string
	Value string
}

type PicklistDefinition struct {
	Name   string
	Values []PicklistValue
}

type ConnectorTemplate struct {
	ObjectDefinitions []*ObjectDefinition
	Actions           []*ActionDefinition
	Picklists         []*PicklistDefinition
}

var typeMap = map[string]string{
	"double":                    "number",
	"float":                     "number",
	"int32":                     "integer",
	"int64":                     "integer",
	"uint32":                    "integer",
	"uint64":                    "integer",
	"sint32":                    "integer",
	"sint64":                    "integer",
	"fixed32":                   "number",
	"fixed64":                   "number",
	"sfixed32":                  "number",
	"sfixed64":                  "number",
	"bool":                      "boolean",
	"google.protobuf.Timestamp": "date_time",
}

//go:embed templates/connector.rb.tmpl
var connectorTmpl string

func findUsedMessages(
	usedMessages map[string]bool,
	usedEnums map[string]bool,
	enums map[string]*gendoc.Enum,
	messages map[string]*gendoc.Message,
	message *gendoc.Message,
) {
	if message == nil {
		return
	}
	usedMessages[message.FullName] = true
	for _, field := range message.Fields {
		if message, ok := messages[field.FullType]; ok {
			if !usedMessages[field.FullType] {
				usedMessages[field.FullType] = true
				findUsedMessages(usedMessages, usedEnums, enums, messages, message)
			}
		} else if _, ok := enums[field.FullType]; ok {
			if !usedEnums[field.FullType] {
				usedEnums[field.FullType] = true
			}
		}
	}
}

func groupActions(actions map[string][]*Method, service *gendoc.Service, method *gendoc.ServiceMethod) {
	if opts, ok := method.Option("s12.protobuf.workato.workato").(*workato.MethodOptionsWorkato); ok {
		if actions[opts.Resource] == nil {
			actions[opts.Resource] = make([]*Method, 0)
		}

		actions[opts.Resource] = append(actions[opts.Resource], &Method{service, method, opts.Action})
	}
}

func getFieldDef(enums map[string]*gendoc.Enum, messages map[string]*gendoc.Message, field *gendoc.MessageField) *FieldDefinition {
	fieldDef := &FieldDefinition{
		Name:  field.Name,
		Label: field.Description,
		Type:  "string",
	}

	if opts, ok := field.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field").(*options.JSONSchema); ok {
		if opts.Default != "" {
			fieldDef.Default = opts.Default
		}
	}

	// Basic Scalar Types
	if t, ok := typeMap[field.FullType]; ok {
		fieldDef.Type = t
		if t == "boolean" {
			fieldDef.ControlType = "checkbox"
		}
	} else if message, ok := messages[field.FullType]; ok {
		fieldDef.Type = "object"
		fieldDef.PropertiesRef = message.FullName
	}

	if enum, ok := enums[field.FullType]; ok {
		fieldDef.ControlType = "select"
		fieldDef.Picklist = fmt.Sprintf("%s_%s", "enum", EscapeActionName(enum.FullName))
		if field.Label == "repeated" {
			fieldDef.ControlType = "multiselect"
		} else {
			if fieldDef.Default == "" {
				fieldDef.Default = enum.Values[0].Name
			}
		}
	} else {
		if field.Label == "repeated" {
			fieldDef.Of = fieldDef.Type
			fieldDef.Type = "array"
		}
	}

	return fieldDef
}

func GenerateWorkatoConnector(template *gendoc.Template, config *Config) ([]byte, error) {
	//return json.Marshal(template)

	messages := make(map[string]*gendoc.Message)
	enums := make(map[string]*gendoc.Enum)
	usedMessages := make(map[string]bool)
	usedEnums := make(map[string]bool)
	actions := make(map[string][]*Method)

	connectorTemplate := ConnectorTemplate{
		ObjectDefinitions: make([]*ObjectDefinition, 0),
		Actions:           make([]*ActionDefinition, 0),
		Picklists:         make([]*PicklistDefinition, 0),
	}

	for _, file := range template.Files {
		for _, message := range file.Messages {
			messages[message.FullName] = message
		}

		for _, enum := range file.Enums {
			enums[enum.FullName] = enum
		}
	}

	for _, file := range template.Files {
		for _, service := range file.Services {
			if service.Name == "WebhooksService" || service.Name == "ThePubService" || service.Name == "InspectionService" || service.Name == "WoraktoService" {
				for _, method := range service.Methods {
					groupActions(actions, service, method)
					findUsedMessages(
						usedMessages,
						usedEnums,
						enums,
						messages,
						messages[method.RequestFullType],
					)
					findUsedMessages(
						usedMessages,
						usedEnums,
						enums,
						messages,
						messages[method.ResponseFullType],
					)
				}
			}
		}
	}

	for msg := range usedMessages {
		message := messages[msg]
		obj := &ObjectDefinition{
			// Use the full name so it is unique
			Name: message.FullName,
		}

		for _, field := range message.Fields {
			obj.Fields = append(obj.Fields, getFieldDef(enums, messages, field))
		}

		connectorTemplate.ObjectDefinitions = append(connectorTemplate.ObjectDefinitions, obj)
	}

	for resource, action := range actions {
		picklistDef := &PicklistDefinition{
			Name:   fmt.Sprintf("%s_%s", "action_name", resource),
			Values: []PicklistValue{},
		}
		actionDef := &ActionDefinition{
			Name:        resource,
			Title:       fmt.Sprintf("Interact with %s %s", indefiniteArticle(resource), resource),
			Subtitle:    fmt.Sprintf("Allows you to intereact with %s %s in iAuditor", indefiniteArticle(resource), resource),
			Description: fmt.Sprintf("<span class='provider'>#{picklist_label['action_name'] || 'Interact with'}</span> %s <span class='provider'>%s</span> in <span class='provider'>iAuditor</span>", indefiniteArticle(resource), resource),
			ConfigFields: []*FieldDefinition{
				{
					Name:        "action_name",
					Label:       "Action",
					Type:        "string",
					ControlType: "select",
					Picklist:    picklistDef.Name,
				},
			},
			InputFields:  make(map[string]string),
			OutputFields: make(map[string]string),
			Endpoints:    make(map[string]Endpoint),
		}

		for _, method := range action {
			name := EscapeActionName(fmt.Sprintf("%s/%s", method.Service.FullName, method.Method.Name))

			actionDef.Endpoints[name] = getExecuteCode(config, messages, method.Service, method.Method)

			picklistDef.Values = append(picklistDef.Values, PicklistValue{name, method.Action})

			actionDef.InputFields[name] = method.Method.RequestFullType
			actionDef.OutputFields[name] = method.Method.ResponseFullType
		}

		// if configActions != nil {
		// 	for _, act := range configActions {

		// 	}
		// }

		connectorTemplate.Actions = append(connectorTemplate.Actions, actionDef)
		connectorTemplate.Picklists = append(connectorTemplate.Picklists, picklistDef)
	}

	for enum := range usedEnums {
		enum := enums[enum]
		pickListDef := &PicklistDefinition{
			Name:   fmt.Sprintf("%s_%s", "enum", EscapeActionName(enum.FullName)),
			Values: []PicklistValue{},
		}

		for _, value := range enum.Values {
			desc := value.Description
			if desc == "" {
				desc = value.Name
			}
			pickListDef.Values = append(pickListDef.Values, PicklistValue{value.Name, desc})
		}

		connectorTemplate.Picklists = append(connectorTemplate.Picklists, pickListDef)
	}

	tp, err := tmpl.New("Connector Template").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(connectorTmpl)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = tp.Execute(&buf, connectorTemplate)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
