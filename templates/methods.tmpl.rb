{{ define "methods" }}
  # Default after_error_response handler, can be overwritten by defining a new method after this one.
  after_error_response: lambda do |code, body, headers, message|
    err = [
      message,
      body,
    ].join("\n\n")

    trace = headers.dig('traceparent') || headers.dig('grpc_metadata_traceparent')
    if trace.present?
      err = "#{err}\n\nTrace ID: #{trace}"
    end

    error(err)
  end,

  "encode_array_to_query_params": lambda do |val|
    val.each do |key, value|
      if (value.is_a? String) && (value[0] == '[') && (value[-1] == ']')
        val[key] = parse_json(value)
      end
    end
    val.encode_www_form
  end,

  # This method is for Custom action
  make_schema_builder_fields_sticky: lambda do |schema|
    schema.map do |field|
      if field['properties'].present?
        field['properties'] = call('make_schema_builder_fields_sticky',
                                   field['properties'])
      end
      field['sticky'] = true

      field
    end
  end,

  # Formats input/output schema to replace any special characters in name,
  # without changing other attributes (method required for custom action)
  format_schema: lambda do |input|
    input&.map do |field|
      if (props = field[:properties])
        field[:properties] = call('format_schema', props)
      elsif (props = field['properties'])
        field['properties'] = call('format_schema', props)
      end
      if (name = field[:name])
        field[:label] = field[:label].presence || name.labelize
        field[:name] = name
                       .gsub(/\W/) { |spl_chr| "__#{spl_chr.encode_hex}__" }
      elsif (name = field['name'])
        field['label'] = field['label'].presence || name.labelize
        field['name'] = name
                        .gsub(/\W/) { |spl_chr| "__#{spl_chr.encode_hex}__" }
      end

      field
    end
  end,

  # Formats payload to inject any special characters that previously removed
  format_payload: lambda do |payload|
    if payload.is_a?(Array)
      payload.map do |array_value|
        call('format_payload', array_value)
      end
    elsif payload.is_a?(Hash)
      payload.each_with_object({}) do |(key, value), hash|
        key = key.gsub(/__[0-9a-fA-F]+__/) do |string|
          string.gsub(/__/, '').decode_hex.as_utf8
        end
        value = call('format_payload', value) if value.is_a?(Array) || value.is_a?(Hash)
        hash[key] = value
      end
    end
  end,

  # Formats response to replace any special characters with valid strings
  # (method required for custom action)
  format_response: lambda do |response|
    response = response&.compact unless response.is_a?(String) || response
    if response.is_a?(Array)
      response.map do |array_value|
        call('format_response', array_value)
      end
    elsif response.is_a?(Hash)
      response.each_with_object({}) do |(key, value), hash|
        key = key.gsub(/\W/) { |spl_chr| "__#{spl_chr.encode_hex}__" }
        if value.is_a?(Array) || value.is_a?(Hash)
          value = call('format_response', value)
        end
        hash[key] = value
      end
    else
      response
    end
  end,
{{- range .Methods }}
  "{{.Name}}": lambda do |{{ .Params | join ", " }}|
    {{- .Exec | nindent 4 }}
  end,
{{- end }}
{{- end }}
