package template

import (
	"fmt"

	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

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

func (t *WorkatoTemplate) generateObjectDefinitions() {
	for _, message := range t.messages {
		obj := &schema.ObjectDefinition{
			// Use the full name so it is unique
			Key: message.FullName,
		}

		if cfg, ok := t.config.Message[message.FullName]; ok && cfg.Exec != "" {
			obj.Exec = cfg.Exec
		}

		for _, field := range message.Fields {
			if opts, ok := field.Option("s12.protobuf.workato.field").(*workato.FieldOptionsWorkato); ok {
				if opts.Excluded {
					continue
				}
			}
			obj.Fields = append(obj.Fields, t.getFieldDef(field))
		}

		t.ObjectDefinitions = append(t.ObjectDefinitions, obj)
	}
}

func (t *WorkatoTemplate) getFieldDef(field *gendoc.MessageField) *schema.FieldDefinition {
	fieldDef := &schema.FieldDefinition{
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
		}
	} else {
		if field.Label == "repeated" {
			fieldDef.Of = fieldDef.Type
			fieldDef.Type = "array"
		}
	}

	if opts, ok := field.Option("s12.protobuf.workato.field").(*workato.FieldOptionsWorkato); ok {
		if opts.DynamicPicklist != "" {
			picklistName := escapeKeyName(opts.DynamicPicklist)
			picklist, ok := t.dynamicPicklistMap[picklistName]
			if !ok {
				panic(fmt.Errorf("invalid dynamic picklist %s for %s", opts.DynamicPicklist, field.Name))
			}

			fieldDef.ControlType = "select"
			fieldDef.Picklist = picklist.Name
			if field.Label == "repeated" {
				fieldDef.ControlType = "multiselect"
			}
		}
	}

	return fieldDef
}
