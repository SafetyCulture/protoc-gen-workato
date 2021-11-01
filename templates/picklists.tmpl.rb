{{ define "picklists" -}}
{{- range . }}
  "{{.Name}}": lambda do
    [
      {{- range .Values }}
      ["{{.Value}}", "{{.Key}}"],
      {{- end }}
    ]
  end,
{{- end }}
{{- end }}
