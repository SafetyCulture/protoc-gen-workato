{{ define "triggers" -}}
{{- range $trigger := . }}
    {{ $trigger.Key }}: {
       title: '{{ $trigger.Value.Title }}',

       description: "{{ $trigger.Value.Description }}",

       input_fields: lambda do |object_definitions|
         object_definitions['{{ $trigger.Value.InputField }}']
       end,

       webhook_subscribe: lambda do |webhook_url, connection, input|
         result = post('/webhooks/v1/webhooks')
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
         if event.has_key?('workflow_id')
           event['workflow_id'] + '@' + event['event']['date_triggered']
         elsif event.has_key?('webhook_id')
           event['webhook_id'] + '@' + event['event']['date_triggered']
         end
       end,

       output_fields: lambda do |object_definitions|
         object_definitions['{{ $trigger.Value.OutputField }}']
       end,
    },
{{- end }}
{{- end }}
