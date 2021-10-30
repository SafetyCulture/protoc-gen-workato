{{ define "field" -}}
{
  name: "{{ .Name }}",
  label: "{{ .Label }}",
  optional: "{{ .Optional }}",
  type: "{{ .Type }}",
  hint: "{{ .Hint }}",
  {{if .Of}}of: "{{ .Of }}",{{end}}
  {{if .PropertiesRef}}properties: object_definitions['{{.PropertiesRef}}'],{{end}}
  {{if .ControlType}}control_type: "{{ .ControlType }}",{{end}}
  #toggle_hint: "{{ .ToggleHint }}",
  #toggle_field: "{{ .ToggleField }}",
  {{if .Default}}default: "{{ .Default }}",{{end}}
  pick_list: "{{ .Picklist }}",
  #delimiter: "{{ .Delimiter }}",
  sticky: true,
  #render_input: "{{ .RenderInput }}",
  #parse_output: "{{ .ParseOutput }}",
  #change_on_blur: "{{ .ChangeOnBlur }}",
  #support_pills: "{{ .SupportPills }}",
  #custom: "{{ .Custom }}",
  #extends_schema: "{{ .ExtendsSchema }}",
  #list_mode: "{{ .ListMode }}",
  #list_mode_toggle: "{{ .ListModeToggle }}",
  #item_label: "{{ .ItemLabel }}",
  #add_field_label: "{{ .AddFieldLabel }}",
  #empty_schema_message: "{{ .EmptySchemaMessage }}",
  #sample_data_type: "{{ .SampleDataType }}",
  #ng_if: "{{ .NgIf }}",
}
{{- end }}
