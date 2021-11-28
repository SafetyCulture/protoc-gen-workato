{{ define "triggers" -}}
{{- range $trigger := . }}
    {{ $trigger.Key }}: {
       title: "{{ $trigger.Value.Title }}",

       description: "{{ $trigger.Value.Description }}",

       input_fields: lambda do |object_definitions|
        {{- range $name, $object_name := $trigger.Value.InputFields }}
        object_definitions["{{ $object_name }}"]
        {{- end }}
       end,

       webhook_subscribe: lambda do |webhook_url, connection, input|
           post("/webhooks/v1/webhooks")
             .payload(
               url: webhook_url,
               trigger_events: ["#{input['trigger']}"]
             )
       end,

       webhook_notification: lambda do |input, payload|
         payload
       end,

       webhook_unsubscribe: lambda do |webhook|
         delete("/webhooks/v1/webhooks/#{webhook['webhook']['webhook_id']}")
       end,

       dedup: lambda do |event|
        event["workflow_id"] + "@" + event["event"]["date_triggered"]
       end,

       output_fields: lambda do |object_definitions|
         {{- range $name, $object_name := $trigger.Value.OutputFields }}
         object_definitions["{{ $object_name }}"]
         {{- end }}
       end,
    },
{{- end }}
{{- end }}
