package extensions

import (
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		opt, ok := payload.(annotations.FieldBehavior)
		if !ok {
			return nil
		}

		return opt
	})
}
