{{ define "methods" -}}
{{- range . }}
  "{{.Name}}": lambda do |{{ .Params | join ", " }}|
    {{- .Exec | nindent 4 }}
  end,
{{- end }}
{{- end }}
