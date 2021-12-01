package template

import (
	"fmt"
)

// TriggerValue is the representation of a trigger value in the Workato SDK
type TriggerValue struct {
	Title       string
	Description string
	InputField  string
	OutputField string
}

// TriggerDefinition is the representation of a trigger in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/triggers.html
type TriggerDefinition struct {
	Key   string
	Value *TriggerValue
}

func (t *WorkatoTemplate) generateTriggerDefinitions() error {
	for _, trigger := range t.triggers {
		tag, err := trigger.extractFirstTag()
		if err != nil {
			return err
		}

		triggerDef := trigger.mapToWorkatoTrigger(tag)
		t.Triggers = append(t.Triggers, triggerDef)
		t.recordUsedTrigger(trigger)
	}

	return nil
}

// recordUsedTrigger registers the usage of the trigger request and response methods in template message Map
func (t *WorkatoTemplate) recordUsedTrigger(trigger *ServiceMethod) {
	t.recordUsedMessage(t.messageMap[trigger.Method.RequestFullType])
	t.recordUsedMessage(t.messageMap[trigger.Method.ResponseFullType])
}

// mapToWorkatoTrigger converts to Workato Format
// It returns pointer to TriggerDefinition
func (t *ServiceMethod) mapToWorkatoTrigger(tag string) *TriggerDefinition {
	triggerDef := TriggerDefinition{
		Key: escapeKeyName("trigger_" + tag),
		Value: &TriggerValue{
			Title:       t.Method.Description,
			Description: fmt.Sprintf("<span class='provider'>Trigger for %s</span>", t.Method.Description),
			InputField:  t.Method.RequestFullType,
			OutputField: t.Method.ResponseFullType,
		},
	}
	return &triggerDef
}
