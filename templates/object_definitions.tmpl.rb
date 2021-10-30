{{ define "object_definitions" -}}
{{- range . }}
  "{{ .Name }}": {
    fields: lambda do |connection, config_fields, object_definitions|
      [
        {{- range .Fields }}
          {{- include "field" . | nindent 8 }},
        {{- end}}
      ]
    end
  },
{{end}}
{{- end }}
