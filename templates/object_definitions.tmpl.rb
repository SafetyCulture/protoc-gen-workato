{{ define "object_definitions" -}}
{{- range . }}
  "{{ .Key }}": {
    fields: lambda do |connection, config_fields, object_definitions|
      {{- if .CustomCode }}
        {{- .CustomCode | nindent 6 }}
      {{- else }}
      [
        {{- range .Fields }}
          {{- include "field" . | nindent 8 }},
        {{- end}}
      ]
      {{- end }}
    end
  },
{{end}}
{{- end }}
