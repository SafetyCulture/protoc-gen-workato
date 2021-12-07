{
  title: "iAuditor by SafetyCulture",

  # API key authentication example. See more examples at https://docs.workato.com/developing-connectors/sdk/guides/authentication.html
  connection: {
    fields: [
      {
        name: 'api_key',
        label: 'API Token',
        optional: false,
        control_type: 'password',
         hint: 'Get your <b>API key</b> <a href="https://app.safetyculture.com/account/api-tokens" target="_blank">here</a>.'
      }
    ],

    authorization: {
      type: 'api_key',

      apply: lambda do |connection|
        headers('Authorization': "Bearer #{connection['api_key']}")
        headers('sc-integration-id': "workato")
        headers('sc-integration-version': "1")
      end
    },

    base_uri: lambda do
      "https://api.safetyculture.io"
    end
  },

  test: lambda do |_connection|
    get('/accounts/user/v1/user:WhoAmI')
  end,

  object_definitions: {
    {{- include "object_definitions" .ObjectDefinitions | indent 2}}
  },

  actions: {
    {{- include "actions" .Actions | indent 2}}
  },

  # Dynamic webhook example. Subscribes and unsubscribes webhooks programmatically
  # see more at https://docs.workato.com/developing-connectors/sdk/guides/building-triggers/dynamic-webhook.html
  triggers: {
    {{- include "triggers" .Triggers | indent 2}}
  },

  pick_lists: {
    {{- include "picklists" .Picklists | indent 2}}
  },

  # Reusable methods can be called from object_definitions, picklists or actions
  # See more at https://docs.workato.com/developing-connectors/sdk/sdk-reference/methods.html
  methods: {
    {{- include "methods" .Methods | indent 2}}
  }
}
