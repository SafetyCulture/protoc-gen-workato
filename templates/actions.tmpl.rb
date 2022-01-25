{{ define "actions" -}}
{{- range $action := . }}
  "{{$action.Name}}": {
    title: "{{ $action.Title }}",
    subtitle: "{{ $action.Subtitle }}",
    description: lambda do |input, picklist_label|
      "{{ $action.Description }}"
    end,
    help: lambda do |input, picklist_label|
      case input['action_name']
      {{- range $name, $message := $action.HelpMessages }}
      when '{{ $name }}'
        {
          body: '{{ $message.Body }}',
          learn_more_url: '{{ $message.LearnMoreURL }}',
          learn_more_text: '{{ $message.LearnMoreText }}'
        }
      {{- end }}
      else
        {
          body: '{{ $action.DefaultHelpMessage.Body }}',
          learn_more_url: '{{ $action.DefaultHelpMessage.LearnMoreURL }}',
          learn_more_text: '{{ $action.DefaultHelpMessage.LearnMoreText }}'
        }
      end
    end,
    config_fields: [
      {{- range $action.ConfigFields }}
        {{- include "field" . | nindent 6 }},
      {{- end }}
    ],
    input_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.InputFields }}
      when '{{ $name }}'
        object_definitions['{{ $object_name }}']
      {{- end }}
      end
    end,
    execute: lambda do |connection, input, eis, eos, continue|
      case input['action_name']
      {{- range $name, $value := $action.ExecCode }}
      when '{{ $name }}'
        exclude_keys = {{ format_string_slice $value.ExcludeFromQuery }}
        body = input.select { |k, v| k != 'action_name' and not exclude_keys.include? k }
        {{ $value.Aggregate | indent 8 | trim }}
      {{- end }}
      end
    end,
    output_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.OutputFields }}
      when '{{ $name }}'
        object_definitions['{{ $object_name }}']
      {{- end }}
      end
    end,
  },
{{end}}
{{- end }}
