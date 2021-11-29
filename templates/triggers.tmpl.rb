{{ define "triggers" -}}
{{- range $trigger := . }}
    {{ $trigger.Key }}: {
       title: "{{ $trigger.Value.Title }}",

       description: "{{ $trigger.Value.Description }}",

       input_fields: lambda do |object_definitions|
        object_definitions["{{ $trigger.Value.InputField }}"]
       end,

       webhook_subscribe: lambda do |webhook_url, connection, input|
           result = post("/webhooks/v1/webhooks")
             .payload(
               url: webhook_url,
               trigger_events: ["#{input['trigger']}"]
             )
           result['webhook']
       end,

       webhook_notification: lambda do |input, payload|
         payload
       end,

       webhook_unsubscribe: lambda do |webhook|
         delete("/webhooks/v1/webhooks/#{webhook['webhook_id']}")
       end,

       dedup: lambda do |event|
        event["workflow_id"] + "@" + event["event"]["date_triggered"]
       end,

       output_fields: lambda do |object_definitions|
        object_definitions["{{ $trigger.Value.OutputField }}"]
       end,
    },
{{- end }}
{{- end }}
