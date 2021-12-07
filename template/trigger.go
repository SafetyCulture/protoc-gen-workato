package template

import (
	"fmt"

	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
)

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
func (t *ServiceMethod) mapToWorkatoTrigger(tag string) *schema.TriggerDefinition {
	triggerDef := schema.TriggerDefinition{
		Key: escapeKeyName("trigger_" + tag),
		Value: &schema.TriggerValue{
			Title:       t.Method.Description,
			Description: fmt.Sprintf("<span class='provider'>Trigger for %s</span>", t.Method.Description),
			InputField:  t.Method.RequestFullType,
			OutputField: t.Method.ResponseFullType,
		},
	}
	return &triggerDef
}
