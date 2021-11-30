package extensions

import (
	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

func init() {
	extensions.SetTransformer("s12.protobuf.workato.trigger", func(payload interface{}) interface{} {
		opt, ok := payload.(*workato.MethodOptionsWorkatoTrigger)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("s12.protobuf.workato.pick_list", func(payload interface{}) interface{} {
		opt, ok := payload.(*workato.MethodOptionsWorkatoPickList)
		if !ok {
			return nil
		}

		return opt
	})
}
