{{ define "triggers" -}}
{{- range $trigger := . }}
    "{{ $trigger.Key }}": {
       title: "{{ $trigger.Value.Title }}",
       description: {{ $trigger.Value.Description }},
       input_fields: "TODO",
    },
{{- end }}
{{- end }}
