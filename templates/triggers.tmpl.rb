{{ define "triggers" -}}
{{- range $trigger := . }}
    "{{ $trigger.Key }}": {
       title: "{{ $trigger.Value.Title }}",

       description: {{ $trigger.Value.Description }},

       input_fields: lambda do |object_definitions|
        {{- range $name, $object_name := $trigger.Value.InputFields }}
        object_definitions["{{ $object_name }}"]
        {{- end }}
       end,

       webhook_subscribe: lambda do |webhook_url, connection, input|
         post("/workflows")
           .payload(
             steps: [
               {
                 url: webhook_url,
                 type: "webhook"
               }
             ],
             trigger_events: input["events"]
           )
       end,

       webhook_notification: lambda do |input, payload|
         payload
       end,

       webhook_unsubscribe: lambda do |webhook|
         delete("/workflows/#{webhook['workflow_id']}")
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
