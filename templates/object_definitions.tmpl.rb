{{ define "object_definitions" }}
  custom_action_input: {
    fields: lambda do |_connection, config_fields|
      verb = config_fields['verb']
      input_schema = parse_json(config_fields.dig('input', 'schema') || '[]')
      data_props =
        input_schema.map do |field|
          if config_fields['request_type'] == 'multipart' &&
            field['binary_content'] == 'true'
            field['type'] = 'object'
            field['properties'] = [
              { name: 'file_content', optional: false },
              {
                name: 'content_type',
                default: 'text/plain',
                sticky: true
              },
              { name: 'original_filename', sticky: true }
            ]
          end
          field
        end
      data_props = call('make_schema_builder_fields_sticky', data_props)
      input_data =
        if input_schema.present?
          if input_schema.dig(0, 'type') == 'array' &&
            input_schema.dig(0, 'details', 'fake_array')
            {
              name: 'data',
              type: 'array',
              of: 'object',
              properties: data_props.dig(0, 'properties')
            }
          else
            { name: 'data', type: 'object', properties: data_props }
          end
        end

      [
        {
          name: 'path',
          hint: 'Base URI is <b>' \
          '{{ .AppBaseURL }}' \
          '</b> - path will be appended to this URI. Use absolute URI to ' \
          'override this base URI.',
          optional: false
        },
        if %w[post put patch].include?(verb)
          {
            name: 'request_type',
            default: 'json',
            sticky: true,
            extends_schema: true,
            control_type: 'select',
            pick_list: [
              ['JSON request body', 'json'],
              ['URL encoded form', 'url_encoded_form'],
              ['Mutipart form', 'multipart'],
              ['Raw request body', 'raw']
            ]
          }
        end,
        {
          name: 'response_type',
          default: 'json',
          sticky: false,
          extends_schema: true,
          control_type: 'select',
          pick_list: [['JSON response', 'json'], ['Raw response', 'raw']]
        },
        if %w[get options delete].include?(verb)
          {
            name: 'input',
            label: 'Request URL parameters',
            sticky: true,
            add_field_label: 'Add URL parameter',
            control_type: 'form-schema-builder',
            type: 'object',
            properties: [
              {
                name: 'schema',
                sticky: input_schema.blank?,
                extends_schema: true
              },
              input_data
            ].compact
          }
        else
          {
            name: 'input',
            label: 'Request body parameters',
            sticky: true,
            type: 'object',
            properties:
              if config_fields['request_type'] == 'raw'
                [{
                  name: 'data',
                  sticky: true,
                  control_type: 'text-area',
                  type: 'string'
                }]
              else
                [
                  {
                    name: 'schema',
                    sticky: input_schema.blank?,
                    extends_schema: true,
                    schema_neutral: true,
                    control_type: 'schema-designer',
                    sample_data_type: 'json_input',
                    custom_properties:
                      if config_fields['request_type'] == 'multipart'
                        [{
                          name: 'binary_content',
                          label: 'File attachment',
                          default: false,
                          optional: true,
                          sticky: true,
                          render_input: 'boolean_conversion',
                          parse_output: 'boolean_conversion',
                          control_type: 'checkbox',
                          type: 'boolean'
                        }]
                      end
                  },
                  input_data
                ].compact
              end
          }
        end,
        {
          name: 'request_headers',
          sticky: false,
          extends_schema: true,
          control_type: 'key_value',
          empty_list_title: 'Does this HTTP request require headers?',
          empty_list_text: 'Refer to the API documentation and add ' \
          'required headers to this HTTP request',
          item_label: 'Header',
          type: 'array',
          of: 'object',
          properties: [{ name: 'key' }, { name: 'value' }]
        },
        unless config_fields['response_type'] == 'raw'
          {
            name: 'output',
            label: 'Response body',
            sticky: true,
            extends_schema: true,
            schema_neutral: true,
            control_type: 'schema-designer',
            sample_data_type: 'json_input'
          }
        end,
        {
          name: 'response_headers',
          sticky: false,
          extends_schema: true,
          schema_neutral: true,
          control_type: 'schema-designer',
          sample_data_type: 'json_input'
        }
      ].compact
    end
  },

  custom_action_output: {
    fields: lambda do |_connection, config_fields|
      response_body = { name: 'body' }

      [
        if config_fields['response_type'] == 'raw'
          response_body
        elsif (output = config_fields['output'])
          output_schema = call('format_schema', parse_json(output))
          if output_schema.dig(0, 'type') == 'array' &&
            output_schema.dig(0, 'details', 'fake_array')
            response_body[:type] = 'array'
            response_body[:properties] = output_schema.dig(0, 'properties')
          else
            response_body[:type] = 'object'
            response_body[:properties] = output_schema
          end

          response_body
        end,
        if (headers = config_fields['response_headers'])
          header_props = parse_json(headers)&.map do |field|
            if field[:name].present?
              field[:name] = field[:name].gsub(/\W/, '_').downcase
            elsif field['name'].present?
              field['name'] = field['name'].gsub(/\W/, '_').downcase
            end
            field
          end

          { name: 'headers', type: 'object', properties: header_props }
        end
      ].compact
    end
  },
{{ range .ObjectDefinitions }}
  '{{ .Key }}': {
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
