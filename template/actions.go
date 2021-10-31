package template

import (
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// Action group is a grouped set of actions with sub actions
type ActionGroup struct {
	Name    string
	Actions []*Action
}

// Action is a combined service and method defintion
type Action struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

// ActionDefinition is the representation of an action in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/actions.html
type ActionDefinition struct {
	Name        string
	Title       string
	Subtitle    string
	Description string

	ConfigFields []*FieldDefinition
	InputFields  map[string]string
	OutputFields map[string]string
	ExecCode     map[string]ExecCode
}

// ExecCode is the code to be run when executing a function
type ExecCode struct {
	// Exclude these fields from the query, because they are passed into the body or as path params
	ExcludeFromQuery []string
	Func             string
}

// Group actions based on their shared resource name
func (t *WorkatoTemplate) groupActions() {
	for _, action := range t.actions {
		// Group methods by their first tag
		var resource string
		if opts, ok := action.Method.Option("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation").(*options.Operation); ok {
			for _, tag := range opts.Tags {
				if tag != "Public" {
					resource = tag
					break
				}
			}
		}
		if resource == "" {
			continue
		}

		actionGroup := t.groupedActionMap[resource]

		if actionGroup == nil {
			actionGroup = &ActionGroup{
				Name:    resource,
				Actions: make([]*Action, 0),
			}
			t.groupedActionMap[resource] = actionGroup
			t.groupedActions = append(t.groupedActions, actionGroup)
		}
		actionGroup.Actions = append(actionGroup.Actions, action)

		t.recordUsedAction(action)
	}
}

func (t *WorkatoTemplate) recordUsedAction(action *Action) {
	t.recordUsedMessage(t.messageMap[action.Method.RequestFullType])
	t.recordUsedMessage(t.messageMap[action.Method.ResponseFullType])
}

func (t *WorkatoTemplate) recordUsedMessage(message *gendoc.Message) {
	if message == nil {
		return
	}

	// Already recorded
	if t.usedMessageMap[message.FullName] != nil {
		return
	}

	t.usedMessageMap[message.FullName] = message
	t.messages = append(t.messages, message)
	for _, field := range message.Fields {
		if subMessage, ok := t.messageMap[field.FullType]; ok {
			t.recordUsedMessage(subMessage)
		} else if _, ok := t.enumMap[field.FullType]; ok {
			if t.usedEnumMap[field.FullType] == nil {
				t.usedEnumMap[field.FullType] = t.enumMap[field.FullType]
				t.enums = append(t.enums, t.enumMap[field.FullType])
			}
		}
	}
}

func (t *WorkatoTemplate) generateActionDefinitions() {
	for _, actionGroup := range t.groupedActions {
		picklistDef := &PicklistDefinition{
			Name:   actionPicklistName(actionGroup.Name),
			Values: []PicklistValue{},
		}
		actionDef := &ActionDefinition{
			Name:        escapeKeyName(actionGroup.Name),
			Title:       actionGroup.Name,
			Subtitle:    fmt.Sprintf("Interact with %s in iAuditor", actionGroup.Name),
			Description: fmt.Sprintf("<span class='provider'>#{picklist_label['action_name'] || 'Interact with %s'}</span> in <span class='provider'>iAuditor</span>", actionGroup.Name),
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
			ExecCode:     make(map[string]ExecCode),
		}

		for _, action := range actionGroup.Actions {
			name := escapeKeyName(fmt.Sprintf("%s/%s", action.Service.FullName, action.Method.Name))

			actionDef.ExecCode[name] = t.getExecuteCode(action)

			picklistDef.Values = append(picklistDef.Values, PicklistValue{name, action.Method.Description})

			actionDef.InputFields[name] = action.Method.RequestFullType
			actionDef.OutputFields[name] = action.Method.ResponseFullType
		}

		t.Actions = append(t.Actions, actionDef)
		t.Picklists = append(t.Picklists, picklistDef)
	}
}
