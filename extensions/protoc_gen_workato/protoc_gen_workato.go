package extensions

import (
	workato "github.com/SafetyCulture/protoc-gen-workato/proto"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

func init() {
	extensions.SetTransformer("s12.protobuf.workato.method", func(payload interface{}) interface{} {
		opt, ok := payload.(*workato.MethodOptionsWorkato)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("s12.protobuf.workato.field", func(payload interface{}) interface{} {
		opt, ok := payload.(*workato.FieldOptionsWorkato)
		if !ok {
			return nil
		}

		return opt
	})
}
