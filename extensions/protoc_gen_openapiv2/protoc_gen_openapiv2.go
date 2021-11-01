package extensions

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

func init() {
	extensions.SetTransformer("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger", func(payload interface{}) interface{} {
		opt, ok := payload.(*options.Swagger)
		if !ok {
			return nil
		}

		return opt
	})
	extensions.SetTransformer("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation", func(payload interface{}) interface{} {
		opt, ok := payload.(*options.Operation)
		if !ok {
			return nil
		}

		return opt
	})
	extensions.SetTransformer("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema", func(payload interface{}) interface{} {
		opt, ok := payload.(*options.Schema)
		if !ok {
			return nil
		}

		return opt
	})
	extensions.SetTransformer("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_tag", func(payload interface{}) interface{} {
		opt, ok := payload.(*options.Tag)
		if !ok {
			return nil
		}

		return opt
	})
	extensions.SetTransformer("grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field", func(payload interface{}) interface{} {
		opt, ok := payload.(*options.JSONSchema)
		if !ok {
			return nil
		}

		return opt
	})
}
