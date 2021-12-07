package schema

// TriggerValue is the representation of a trigger value in the Workato SDK
type TriggerValue struct {
	Title       string
	Description string
	InputField  string
	OutputField string
}

// TriggerDefinition is the representation of a trigger in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/triggers.html
type TriggerDefinition struct {
	Key   string
	Value *TriggerValue
}
