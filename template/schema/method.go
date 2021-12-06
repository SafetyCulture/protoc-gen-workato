package schema

// MethodDefinition represents a method in the workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/methods.html
type MethodDefinition struct {
	Name   string
	Params []string
	Code   string
}
