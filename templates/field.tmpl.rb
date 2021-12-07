{{ define "field" -}}
{
  {{ field_key "name" .Name -}}
  {{ field_key "label" .Label -}}
  {{ field_key "optional" .Optional -}}
  {{ field_key "type" .Type -}}
  {{ field_key "hint" .Hint -}}
  {{ field_key "of" .Of -}}
  {{ if .PropertiesRef }}properties: object_definitions['{{.PropertiesRef}}'],{{ end }}
  {{ field_key "control_type" .ControlType -}}
  {{ field_key "toggle_hint" .ToggleHint -}}
  {{ if .ToggleField }}toggle_field: {{ include "field" .ToggleField | indent 2 }},{{ end }}
  {{ field_key "default" .Default -}}
  {{ field_key "pick_list" .Picklist -}}
  {{ field_key "delimiter" .Delimiter -}}
  sticky: true,
  {{ field_key "render_input" .RenderInput -}}
  {{ field_key "parse_output" .ParseOutput -}}
  {{ field_key "change_on_blur" .ChangeOnBlur -}}
  {{ field_key "support_pills" .SupportPills -}}
  {{ field_key "custom" .Custom -}}
  {{ field_key "extends_schema" .ExtendsSchema -}}
  {{ field_key "list_mode" .ListMode -}}
  {{ field_key "list_mode_toggle" .ListModeToggle -}}
  {{ field_key "item_label" .ItemLabel -}}
  {{ field_key "add_field_label" .AddFieldLabel -}}
  {{ field_key "empty_schema_message" .EmptySchemaMessage -}}
  {{ field_key "sample_data_type" .SampleDataType -}}
  {{ field_key "ngIf" .NgIf }}
}
{{- end }}
