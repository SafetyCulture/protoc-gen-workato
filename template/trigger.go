package template

import (
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// Trigger is a combined service and method definition
type Trigger struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

// TriggerValue is the representation of a trigger value in the Workato SDK
type TriggerValue struct {
	Title        string
	Description  string
	InputFields  map[string]string
	OutputFields map[string]string
}

// TriggerDefinition is the representation of a trigger in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/triggers.html
type TriggerDefinition struct {
	Key   string
	Value *TriggerValue
}

func (t *WorkatoTemplate) generateTriggerDefinitions() error {
	for _, trigger := range t.triggers {
		tag, err := trigger.ExtractFirstTag()
		if err != nil {
			return err
		}

		triggerDef := trigger.MapToWorkatoFormat(tag)
		t.Triggers = append(t.Triggers, triggerDef)
		t.recordUsedTrigger(trigger)
	}

	return nil
}

// recordUsedTrigger registers the usage of the trigger request and response methods in template message Map
func (t *WorkatoTemplate) recordUsedTrigger(trigger *Trigger) {
	t.recordUsedMessage(t.messageMap[trigger.Method.RequestFullType])
	t.recordUsedMessage(t.messageMap[trigger.Method.ResponseFullType])
}

// ExtractFirstTag Extract and Converts first non-public tag
func (t *Trigger) ExtractFirstTag() (string, error) {
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

	return escapeKeyName(tagName), nil
}

// MapToWorkatoFormat converts to Workato Format
// It returns pointer to TriggerDefinition
func (t *Trigger) MapToWorkatoFormat(tag string) *TriggerDefinition {
	triggerDef := TriggerDefinition{
		Key: tag,
		Value: &TriggerValue{
			Title:        t.Method.Description,
			Description:  fmt.Sprintf("<span class='provider'>Trigger for %s</span>", t.Method.Description),
			InputFields:  make(map[string]string),
			OutputFields: make(map[string]string),
		},
	}

	name := escapeKeyName(fmt.Sprintf("%s/%s", t.Service.FullName, t.Method.Name))
	triggerDef.Value.InputFields[name] = t.Method.RequestFullType
	triggerDef.Value.OutputFields[name] = t.Method.ResponseFullType
	return &triggerDef
}
