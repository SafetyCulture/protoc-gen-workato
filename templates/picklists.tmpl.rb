{{ define "picklists" -}}
{{- range . }}
  "{{.Name}}": lambda do
    {{- if .Exec }}
      {{- .Exec | indent 4 }}
    {{- else }}
    [
      {{- range .Values }}
      ['{{.Value}}', '{{.Key}}'],
      {{- end }}
    ]
    {{- end }}
  end,
{{- end }}
{{- end }}
