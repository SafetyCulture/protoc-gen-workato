package template

type PicklistValue struct {
	Key   string
	Value string
}

type PicklistDefinition struct {
	Name   string
	Values []PicklistValue
}
