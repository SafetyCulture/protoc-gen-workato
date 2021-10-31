package template

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
}
