package template

import (
	"fmt"
	"strings"

	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// PicklistValue is the value of a picklist item
type PicklistValue struct {
	Key   string
	Value string
}

// PicklistDefinition is the definition of a picklist
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/picklists.html
type PicklistDefinition struct {
	Name   string
	Values []PicklistValue
	Exec   string
}

func (t *WorkatoTemplate) generateDynamicPickList(service *gendoc.Service, method *gendoc.ServiceMethod) *PicklistDefinition {
	var opt *workato.MethodOptionsWorkatoPickList
	var ok bool
	if opt, ok = method.Option("s12.protobuf.workato.pick_list").(*workato.MethodOptionsWorkatoPickList); !ok {
		return nil
	}

	actionCode := t.getExecuteCode(service, method)

	labelPath := strings.Split(opt.Label, ".")
	valuePath := strings.Split(opt.Value, ".")

	listPath := ""
	for i, value := range valuePath {
		if i == len(valuePath)-1 {
			break
		}
		listPath = fmt.Sprintf("%s['%s']", listPath, value)
	}

	if len(labelPath) != len(valuePath) {
		panic(fmt.Errorf("%s/%s: s12.protobuf.workato.pick_list label and value path not equal depth", service.FullName, method.Name))
	}

	execCode := fmt.Sprintf(`
body = {}
resp = %s

resp%s.pluck('%s', '%s')`,
		actionCode.Func, listPath, labelPath[len(labelPath)-1], valuePath[len(valuePath)-1])

	return &PicklistDefinition{
		Name: dynamicPickListName(fullActionName(service, method)),
		Exec: execCode,
	}
}
