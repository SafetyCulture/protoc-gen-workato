package template

import (
	"fmt"
	"strings"

	workato "github.com/SafetyCulture/protoc-gen-workato/s12/protobuf/workato"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
)

func (t *WorkatoTemplate) recordDynamicPicklist(serviceMethod *ServiceMethod, opt *workato.MethodOptionsWorkato) *schema.PicklistDefinition {
	service := serviceMethod.Service
	method := serviceMethod.Method

	actionCode := t.getExecuteCode(service, method)

	labelPath := strings.Split(opt.Picklist.Label, ".")
	valuePath := strings.Split(opt.Picklist.Value, ".")

	listPath := ""
	for i, value := range valuePath {
		if i == len(valuePath)-1 {
			break
		}
		listPath = fmt.Sprintf("%s['%s']", listPath, value)
	}

	if len(labelPath) != len(valuePath) {
		panic(fmt.Errorf("%s/%s: s12.protobuf.workato.field.picklist label and value path not equal depth", service.FullName, method.Name))
	}

	execCode := fmt.Sprintf(`
body = {}
%s
resp = %s
resp%s.pluck('%s', '%s')`,
		actionCode.Body, actionCode.Func, listPath, labelPath[len(labelPath)-1], valuePath[len(valuePath)-1])

	name := fullActionName(service, method)

	picklist := &schema.PicklistDefinition{
		Name: dynamicPicklistName(name),
		Exec: execCode,
	}

	t.dynamicPicklistMap[name] = picklist

	return picklist
}
