{{ define "actions" -}}
{{- range $action := . }}
  "{{$action.Name}}": {
    title: "{{ $action.Title }}",
    subtitle: "{{ $action.Subtitle }}",
    description: lambda do |input, picklist_label|
      "{{ $action.Description }}"
    end,
    config_fields: [
      {{- range $action.ConfigFields }}
        {{- include "field" . | nindent 6 }},
      {{- end }}
    ],
    input_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.InputFields }}
      when "{{ $name }}"
        object_definitions["{{ $object_name }}"]
      {{- end }}
      end
    end,
    execute: lambda do |connection, input, eis, eos, continue|
      case input['action_name']
      {{- range $name, $value := $action.ExecCode }}
      when "{{ $name }}"
        excludeKeys = {{ format_string_slice $value.ExcludeFromQuery }}
        body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
        {{ $value.Func | indent 8 | trim }}
      {{- end }}
      end
    end,
    output_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.OutputFields }}
      when "{{ $name }}"
        object_definitions["{{ $object_name }}"]
      {{- end }}
      end
    end,
  },
{{end}}
{{- end }}