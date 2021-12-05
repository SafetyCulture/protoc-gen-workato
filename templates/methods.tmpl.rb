{{ define "methods" -}}
{{- range . }}
  "{{.Name}}": lambda do |{{ .Params | join ", " }}|
    {{- .Code | nindent 4 }}
  end,
{{- end }}
{{- end }}
