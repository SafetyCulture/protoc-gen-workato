package template

import (
	"fmt"

	workato "github.com/SafetyCulture/protoc-gen-workato/s12/protobuf/workato"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"google.golang.org/genproto/googleapis/api/annotations"
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

// capture messages that are manually included via the config
func (t *WorkatoTemplate) captureIncludedMessages() {
	for msg, cfg := range t.config.Message {
		if message, ok := t.messageMap[msg]; cfg.Include && ok {
			t.recordUsedMessage(message)
		}
	}
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
			if !t.checkVisibility(field.Option("google.api.field_visibility")) {
				continue
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
		Hint:  markdownToHTML(field.Description),
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
		switch fieldType {
		case "boolean":
			fieldDef.ControlType = "checkbox"
			fieldDef.ConvertInput = "boolean_conversion"
		case "integer":
			fieldDef.ConvertInput = "integer_conversion"
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
		if opts.DynamicPicklist != "" || opts.Picklist != "" {
			picklistName := opts.Picklist
			if opts.DynamicPicklist != "" {
				picklistName = escapeKeyName(opts.DynamicPicklist)
				picklist, ok := t.dynamicPicklistMap[picklistName]
				if !ok {
					panic(fmt.Errorf("invalid dynamic picklist %s for %s", opts.DynamicPicklist, field.Name))
				}
				// This name is prefixed, so we need to use the right one.
				picklistName = picklist.Name
			}

			fieldDef.ControlType = "select"
			fieldDef.Picklist = picklistName
			fieldDef.ToggleHint = "Select from list"
			if field.Label == "repeated" {
				fieldDef.ControlType = "multiselect"
			}

			toggleFieldDef := *fieldDef
			toggleFieldDef.ControlType = "text"
			toggleFieldDef.Picklist = ""
			toggleFieldDef.ToggleHint = "Use ID"

			fieldDef.ToggleField = &toggleFieldDef
		}
	}

	if fieldBehavior, ok := field.Option("google.api.field_behavior").([]annotations.FieldBehavior); ok {
		for _, behaviour := range fieldBehavior {
			switch behaviour {
			case annotations.FieldBehavior_REQUIRED:
				fieldDef.Optional = boolPtr(false)
			default:
			}
		}
	}

	return fieldDef
}
