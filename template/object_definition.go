package template

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

type ObjectDefinition struct {
	Name string

	Fields []*FieldDefinition
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

func (t *WorkatoTemplate) generateObjectDefintions() {
	for _, message := range t.messages {
		obj := &ObjectDefinition{
			// Use the full name so it is unique
			Name: message.FullName,
		}

		for _, field := range message.Fields {
			obj.Fields = append(obj.Fields, t.getFieldDef(field))
		}

		t.ObjectDefinitions = append(t.ObjectDefinitions, obj)
	}
}

func (t *WorkatoTemplate) getFieldDef(field *gendoc.MessageField) *FieldDefinition {
	fieldDef := &FieldDefinition{
		Name:  field.Name,
		Label: fieldTitleFromName(field.Name),
		Hint:  field.Description,
		Type:  "string",
	}

	if opts, ok := field.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field").(*options.JSONSchema); ok {
		if opts.Title != "" {
			fieldDef.Label = opts.Title
		}

		if opts.Description != "" {
			fieldDef.Hint = opts.Description
		}

		if opts.Default != "" {
			fieldDef.Default = opts.Default
		}
	}

	// Basic Scalar Types
	if fieldType, ok := typeMap[field.FullType]; ok {
		fieldDef.Type = fieldType
		if fieldType == "boolean" {
			fieldDef.ControlType = "checkbox"
		}
	} else if message, ok := t.messageMap[field.FullType]; ok {
		fieldDef.Type = "object"
		fieldDef.PropertiesRef = message.FullName
	}

	if enum, ok := t.enumMap[field.FullType]; ok {
		fieldDef.ControlType = "select"
		fieldDef.Picklist = enumPicklistName(enum)
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
