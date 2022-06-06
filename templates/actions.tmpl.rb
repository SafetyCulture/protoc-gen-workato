{{ define "actions" }}
  custom_action: {
    title: 'Custom Action',
    subtitle: 'Build your own {{ .Name }} action with a HTTP request',

    description: lambda do |object_value, _object_label|
      "<span class='provider'>" \
      "#{object_value[:action_name] || 'Custom action'}</span> in " \
      "<span class='provider'>{{ .Name }}</span>"
    end,

    help: {
      body: 'Build your own {{ .Name }} action with a HTTP request. ' \
      'The request will be authorized with your {{ .Name }} connection.',
      learn_more_url: '{{ .DeveloperDocsURL }}',
      learn_more_text: '{{ .Name }} API documentation'
    },

    config_fields: [
      {
        name: 'action_name',
        hint: "Give this action you're building a descriptive name, e.g. " \
        'create record, get record',
        default: 'Custom action',
        optional: false,
        schema_neutral: true
      },
      {
        name: 'verb',
        label: 'Method',
        hint: 'Select HTTP method of the request',
        optional: false,
        control_type: 'select',
        pick_list: %w[get post put patch options delete]
          .map { |verb| [verb.upcase, verb] }
      }
    ],

    input_fields: lambda do |object_definition|
      object_definition['custom_action_input']
    end,

    execute: lambda do |_connection, input|
      verb = input['verb']
      if %w[get post put patch options delete].exclude?(verb)
        error("#{verb.upcase} not supported")
      end
      path = input['path']
      data = input.dig('input', 'data') || {}
      if input['request_type'] == 'multipart'
        data = data.each_with_object({}) do |(key, val), hash|
          hash[key] = if val.is_a?(Hash)
                        [val[:file_content],
                        val[:content_type],
                        val[:original_filename]]
                      else
                        val
                      end
        end
      end
      request_headers = input['request_headers']
        &.each_with_object({}) do |item, hash|
        hash[item['key']] = item['value']
      end || {}
      request = case verb
                when 'get'
                  get(path, data)
                when 'post'
                  if input['request_type'] == 'raw'
                    post(path).request_body(data)
                  else
                    post(path, data)
                  end
                when 'put'
                  if input['request_type'] == 'raw'
                    put(path).request_body(data)
                  else
                    put(path, data)
                  end
                when 'patch'
                  if input['request_type'] == 'raw'
                    patch(path).request_body(data)
                  else
                    patch(path, data)
                  end
                when 'options'
                  options(path, data)
                when 'delete'
                  delete(path, data)
                end.headers(request_headers)
      request = case input['request_type']
                when 'url_encoded_form'
                  request.request_format_www_form_urlencoded
                when 'multipart'
                  request.request_format_multipart_form
                else
                  request
                end
      response =
        if input['response_type'] == 'raw'
          request.response_format_raw
        else
          request
        end
        .after_error_response(/.*/) do |code, body, headers, message|
          error({ code: code, message: message, body: body, headers: headers }
            .to_json)
        end

      response.after_response do |_code, res_body, res_headers|
        {
          body: res_body ? call('format_response', res_body) : nil,
          headers: res_headers
        }
      end
    end,

    output_fields: lambda do |object_definition|
      object_definition['custom_action_output']
    end,
    retry_on_response: [429, 500, 502, 503, 504, 507, 524],
    max_retries: 3,
  },
{{ range $action := .Actions }}
  "{{$action.Name}}": {
    title: "{{ $action.Title }}",
    subtitle: "{{ $action.Subtitle }}",
    description: lambda do |input, picklist_label|
      "{{ $action.Description }}"
    end,
    help: lambda do |input, picklist_label|
      case input['action_name']
      {{- range $name, $message := $action.HelpMessages }}
      when '{{ $name }}'
        {
          body: '{{ $message.Body }}',
          learn_more_url: '{{ $message.LearnMoreURL }}',
          learn_more_text: '{{ $message.LearnMoreText }}'
        }
      {{- end }}
      else
        {
          body: '{{ $action.DefaultHelpMessage.Body }}',
          learn_more_url: '{{ $action.DefaultHelpMessage.LearnMoreURL }}',
          learn_more_text: '{{ $action.DefaultHelpMessage.LearnMoreText }}'
        }
      end
    end,
    config_fields: [
      {{- range $action.ConfigFields }}
        {{- include "field" . | nindent 6 }},
      {{- end }}
    ],
    input_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.InputFields }}
      when '{{ $name }}'
        object_definitions['{{ $object_name }}']
      {{- end }}
      end
    end,
    execute: lambda do |connection, input, eis, eos, continue|
      case input['action_name']
      {{- range $name, $value := $action.ExecCode }}
      when '{{ $name }}'
        exclude_keys = {{ format_string_slice $value.ExcludeFromQuery }}
        body = input.select { |k, v| k != 'action_name' and not exclude_keys.include? k }
        {{ $value.Aggregate | indent 8 | trim }}
      {{- end }}
      end
    end,
    output_fields: lambda do |object_definitions, connection, config_fields|
      case config_fields['action_name']
      {{- range $name, $object_name := $action.OutputFields }}
      when '{{ $name }}'
        object_definitions['{{ $object_name }}']
      {{- end }}
      end
    end,
    retry_on_response: [429, 500, 502, 503, 504, 507, 524],
    max_retries: 3,
  },
{{end}}
{{- end }}
