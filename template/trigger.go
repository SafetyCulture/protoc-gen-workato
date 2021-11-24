package template

import (
	"fmt"
	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"log"
	"strings"
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

func (t *WorkatoTemplate) generateTriggerDefinitions() {
	for _, trigger := range t.triggers {
		opt, err := trigger.ExtractTriggerOption()
		if err != nil {
			log.Fatal(err)
		}

		triggerDef := &TriggerDefinition{
			Key: opt.Resource,
			Value: &TriggerValue{
				Title:        trigger.Method.Name, //TODO INTG-1991 ?
				Description:  fmt.Sprintf("<span class='provider'>%s</span>", trigger.Method.Description),
				InputFields:  make(map[string]string),
				OutputFields: make(map[string]string),
				// TODO INTG-1991. the other fields
			},
		}

		name := escapeKeyName(fmt.Sprintf("%s/%s", trigger.Service.FullName, trigger.Method.Name))
		triggerDef.Value.InputFields[name] = trigger.Method.RequestFullType
		triggerDef.Value.OutputFields[name] = trigger.Method.ResponseFullType
		t.Triggers = append(t.Triggers, triggerDef)
		t.recordUsedTrigger(trigger)
	}
}

func (t *WorkatoTemplate) recordUsedTrigger(trigger *Trigger) {
	t.recordUsedMessage(t.messageMap[trigger.Method.RequestFullType])
	t.recordUsedMessage(t.messageMap[trigger.Method.ResponseFullType])
}

// ExtractTriggerOption Extract and Validates MethodOptionsWorkatoTrigger
// It returns a pair of MethodOptionsWorkatoTrigger and Error
func (t *Trigger) ExtractTriggerOption() (*workato.MethodOptionsWorkatoTrigger, error) {
	res, ok := t.Method.Option("s12.protobuf.workato.trigger").(*workato.MethodOptionsWorkatoTrigger)
	if ok == false {
		return nil, fmt.Errorf("cannot extract s12.protobuf.workato.trigger from method %s", t.Method.Name)
	}

	if valid := isTriggerOptionValid(res); valid == false {
		return nil, fmt.Errorf("the method %s is not valid", t.Method.Name)
	}

	return res, nil
}

// isTriggerOptionValid validates attributes of MethodOptionsWorkatoTrigger
// It returns true if is valid, false if is not valid
func isTriggerOptionValid(t *workato.MethodOptionsWorkatoTrigger) bool {
	if len(strings.TrimSpace(t.Resource)) > 0 {
		return true
	}
	return false
}
