package schema

// ObjectDefinition is the representation of an object in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/object_definitions.html
type ObjectDefinition struct {
	// Object key
	Key string

	// Fields for the object
	Fields []*FieldDefinition

	// Custom code, if set `Fields` will be ignored
	Exec string
}

// FieldDefinition is the representation of an objects fields in the Workato SDK
// https://docs.workato.com/developing-connectors/sdk/sdk-reference/schema.html
type FieldDefinition struct {
	Name               string             `yaml:"name"`
	Label              string             `yaml:"label"`
	Optional           *bool              `yaml:"optional"`
	Type               string             `yaml:"type"`
	Hint               string             `yaml:"hint"`
	Of                 string             `yaml:"of"`
	PropertiesRef      string             `yaml:"properties_ref"`
	Properties         []*FieldDefinition `yaml:"properties"`
	ControlType        string             `yaml:"control_type"`
	ToggleHint         string             `yaml:"toggle_hint"`
	ToggleField        *FieldDefinition   `yaml:"toggle_field"`
	Default            string             `yaml:"default"`
	Picklist           string             `yaml:"picklist"`
	Delimiter          string             `yaml:"delimiter"`
	Sticky             *bool              `yaml:"sticky"`
	RenderInput        string             `yaml:"render_input"`
	ParseOutput        string             `yaml:"parse_output"`
	ChangeOnBlur       *bool              `yaml:"change_on_blur"`
	SupportPills       *bool              `yaml:"support_pills"`
	Custom             *bool              `yaml:"custom"`
	ExtendsSchema      *bool              `yaml:"extends_schema"`
	ListMode           string             `yaml:"list_mode"`
	ListModeToggle     *bool              `yaml:"list_mode_toggle"`
	ItemLabel          string             `yaml:"item_label"`
	AddFieldLabel      string             `yaml:"add_field_label"`
	EmptySchemaMessage string             `yaml:"empty_schema_message"`
	SampleDataType     string             `yaml:"sample_data_type"`
	NgIf               string             `yaml:"ng_if"`
	ConvertInput       string             `yaml:"convert_input"`
}
