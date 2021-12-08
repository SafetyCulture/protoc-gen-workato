{{ define "object_definitions" -}}
{{- range . }}
  "{{ .Key }}": {
    fields: lambda do |connection, config_fields, object_definitions|
      definition = [
        {{- range .Fields }}
          {{- include "field" . | nindent 8 }},
        {{- end}}
      ]

      {{- if .Exec }}
        {{- .Exec | nindent 6 }}
      {{- else }}
      definition
      {{- end }}
    end
  },
{{end}}
{{- end }}
