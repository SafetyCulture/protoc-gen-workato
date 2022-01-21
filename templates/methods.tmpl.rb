{{ define "methods" }}
  "encode_array_to_query_params": lambda do |val|
    val.each do |key, value|
      if value.is_a? String && value[0] == '[' and value[-1] == ']'
        val[key] = parse_json(value)
      end
    end
    val.encode_www_form
  end,
{{- range . }}
  "{{.Name}}": lambda do |{{ .Params | join ", " }}|
    {{- .Exec | nindent 4 }}
  end,
{{- end }}
{{- end }}
