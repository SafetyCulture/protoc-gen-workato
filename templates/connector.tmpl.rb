{
  title: "{{.Name}}",

  # API key authentication example. See more examples at https://docs.workato.com/developing-connectors/sdk/guides/authentication.html
  connection: {
    fields: [
      {
        name: 'api_key',
        label: 'API Key',
        optional: false,
        control_type: 'password',
        hint: 'Get your <b>API key</b> <a href="https://app.example.com/account/api-tokens" target="_blank">here</a>.'
      }
    ],

    authorization: {
      type: 'api_key',

      apply: lambda do |connection|
        headers('Authorization': "Bearer #{connection['api_key']}")
      end
    },

    base_uri: lambda do
      "https://api.example.com"
    end
  },

  test: lambda do |_connection|
    get('/users/me')
  end,

  object_definitions: {
    {{- include "object_definitions" . | indent 2}}
  },

  actions: {
    {{- include "actions" . | indent 2}}
  },

  # Dynamic webhook example. Subscribes and unsubscribes webhooks programmatically
  # see more at https://docs.workato.com/developing-connectors/sdk/guides/building-triggers/dynamic-webhook.html
  triggers: {
    {{- include "triggers" . | indent 2}}
  },

  pick_lists: {
    {{- include "picklists" . | indent 2}}
  },

  # Reusable methods can be called from object_definitions, picklists or actions
  # See more at https://docs.workato.com/developing-connectors/sdk/sdk-reference/methods.html
  methods: {
    {{- include "methods" . | indent 2}}
  }
}
