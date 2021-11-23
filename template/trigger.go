package template

import (
	"fmt"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// Trigger is a combined service and method definition
type Trigger struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

// TriggerValue is the representation of a trigger value in the Workato SDK
type TriggerValue struct {
	Title       string
	Description string
	InputFields map[string]string
}

// TriggerDefinition is the representation of a trigger in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/triggers.html
type TriggerDefinition struct {
	Key   string
	Value *TriggerValue
}

func (t *WorkatoTemplate) generateTriggerDefinitions() {
	for _, trigger := range t.triggers {
		triggerDef := &TriggerDefinition{
			Key: "__KEY", //TODO: SHOULD COME FROM PROTO OPTIONS .resource !!!
			Value: &TriggerValue{
				Title:       trigger.Method.Name, //TODO ?
				Description: fmt.Sprintf("<span class='provider'>%s</span>", trigger.Method.Description),
				InputFields: make(map[string]string),
			},
		}

		name := escapeKeyName(fmt.Sprintf("%s/%s", trigger.Service.FullName, trigger.Method.Name))
		triggerDef.Value.InputFields[name] = trigger.Method.ResponseFullType

		t.Triggers = append(t.Triggers, triggerDef)
	}
}
