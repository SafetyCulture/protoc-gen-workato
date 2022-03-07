package extensions

import (
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/visibility"
)

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		opt, ok := payload.([]annotations.FieldBehavior)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.enum_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.value_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.field_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.message_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.method_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})

	extensions.SetTransformer("google.api.api_visibility", func(payload interface{}) interface{} {
		opt, ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return opt
	})
}
