{{ define "object_definitions" -}}
{{- range . }}
  "{{ .Key }}": {
    fields: lambda do |connection, config_fields, object_definitions|
      {{- if .Exec }}
        {{- .Exec | nindent 6 }}
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
