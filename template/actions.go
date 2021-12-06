package template

import (
	"fmt"

	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// ActionGroup is a grouped set of actions with sub actions
type ActionGroup struct {
	Name    string
	Actions []*ServiceMethod
}

// Group actions based on their shared resource name
func (t *WorkatoTemplate) groupActions() {
	for _, action := range t.actions {
		// Group methods by their first tag
		resource, err := action.extractFirstTag()
		if err != nil {
			continue
		}

		actionGroup := t.groupedActionMap[resource]

		if actionGroup == nil {
			actionGroup = &ActionGroup{
				Name:    resource,
				Actions: make([]*ServiceMethod, 0),
			}
			t.groupedActionMap[resource] = actionGroup
			t.groupedActions = append(t.groupedActions, actionGroup)
		}
		actionGroup.Actions = append(actionGroup.Actions, action)

		t.recordUsedAction(action)
	}
}

func (t *WorkatoTemplate) recordUsedAction(action *ServiceMethod) {
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
		picklistDef := &schema.PicklistDefinition{
			Name:   actionPicklistName(actionGroup.Name),
			Values: []schema.PicklistValue{},
		}
		actionDef := &schema.ActionDefinition{
			Name:        "action_" + escapeKeyName(actionGroup.Name),
			Title:       actionGroup.Name,
			Subtitle:    fmt.Sprintf("Interact with %s in iAuditor", actionGroup.Name),
			Description: fmt.Sprintf("<span class='provider'>#{picklist_label['action_name'] || 'Interact with %s'}</span> in <span class='provider'>iAuditor</span>", actionGroup.Name),
			ConfigFields: []*schema.FieldDefinition{
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
			ExecCode:     make(map[string]schema.ExecCode),
		}

		if cfg, ok := t.config.Action[actionGroup.Name]; ok {
			for _, field := range cfg.InputFields {
				inputField := field

				if inputField.Picklist != "" {
					if picklist, ok := t.dynamicPicklistMap[escapeKeyName(inputField.Picklist)]; ok {
						inputField.Picklist = picklist.Name
					}
				}
				actionDef.ConfigFields = append(actionDef.ConfigFields, &inputField)
			}
		}

		for _, action := range actionGroup.Actions {
			name := fullActionName(action.Service, action.Method)

			actionDef.ExecCode[name] = t.getExecuteCode(action.Service, action.Method)

			picklistDef.Values = append(picklistDef.Values, schema.PicklistValue{
				Key:   name,
				Value: action.Method.Description,
			})

			actionDef.InputFields[name] = action.Method.RequestFullType
			actionDef.OutputFields[name] = action.Method.ResponseFullType
		}

		t.Actions = append(t.Actions, actionDef)
		t.Picklists = append(t.Picklists, picklistDef)
	}
}
