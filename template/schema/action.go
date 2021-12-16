package schema

import "fmt"

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
	Body             string
	Func             string
}

// Aggregate combines Body and Func on 2 separate lines
func (e ExecCode) Aggregate() string {
	return fmt.Sprintf("%s\n%s", e.Body, e.Func)
}
