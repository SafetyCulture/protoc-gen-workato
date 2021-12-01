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
    "api.tasks.v1.CreateTaskRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "name",
            label: "Name",
            type: "string",
            hint: "The name of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.CreateTaskResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.GetTaskRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.GetTaskResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "task",
            label: "Task",
            type: "object",
            properties: object_definitions['api.tasks.v1.Task'],
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.Task": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
          {
            name: "name",
            label: "Name",
            type: "string",
            hint: "The name of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.UpdateTaskRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            type: "string",
            hint: "The ID of the task",
            
            control_type: "select",
            pick_list: "dynamic_api_tasks_v1_tasksservice_listtasks",
            sticky: true,
            
          },
          {
            name: "name",
            label: "Name",
            type: "string",
            hint: "The name of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.UpdateTaskResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
        ]
      end
    },
  
    "api.tasks.v1.DeleteTaskRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.DeleteTaskResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
        ]
      end
    },
  
    "api.tasks.v1.AddCommentRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "task_id",
            label: "Task Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
          {
            name: "comment",
            label: "Comment",
            type: "string",
            hint: "The comment to add",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.AddCommentResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "comment_id",
            label: "Comment Id",
            type: "string",
            hint: "The ID of the comment",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.UpdateCommentRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "task_id",
            label: "Task Id",
            type: "string",
            hint: "The ID of the task",
            
            sticky: true,
            
          },
          {
            name: "comment_id",
            label: "Comment Id",
            type: "string",
            hint: "The ID comment to update",
            
            sticky: true,
            
          },
          {
            name: "comment",
            label: "Comment",
            type: "string",
            hint: "The updated comment",
            
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.UpdateCommentResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
        ]
      end
    },
  
    "api.tasks.v1.TriggerInspectionRequest": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "trigger",
            label: "Trigger",
            type: "string",
            hint: "Trigger event to subscribe to.",
            
            control_type: "select",
            default: "TRIGGER_EVENT_INSPECTION_STARTED",
            pick_list: "enum_api_tasks_v1_triggerinspectionrequest_triggerevent",
            sticky: true,
            
          },
        ]
      end
    },
  
    "api.tasks.v1.TriggerInspectionResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "workflow_id",
            label: "Workflow Id",
            type: "string",
            
            sticky: true,
            
          },
        ]
      end
    },
  
  },

  actions: {  
    "action_tasks": {
      title: "Tasks",
      subtitle: "Interact with Tasks in iAuditor",
      description: lambda do |input, picklist_label|
        "<span class='provider'>#{picklist_label['action_name'] || 'Interact with Tasks'}</span> in <span class='provider'>iAuditor</span>"
      end,
      config_fields: [
        {
          name: "action_name",
          label: "Action",
          type: "string",
          
          control_type: "select",
          pick_list: "action_name_tasks",
          sticky: true,
          
        },
      ],
      input_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_tasksservice_createtask"
          object_definitions["api.tasks.v1.CreateTaskRequest"]
        when "api_tasks_v1_tasksservice_deletetask"
          object_definitions["api.tasks.v1.DeleteTaskRequest"]
        when "api_tasks_v1_tasksservice_gettask"
          object_definitions["api.tasks.v1.GetTaskRequest"]
        when "api_tasks_v1_tasksservice_updatetask"
          object_definitions["api.tasks.v1.UpdateTaskRequest"]
        end
      end,
      execute: lambda do |connection, input, eis, eos, continue|
        case input['action_name']
        when "api_tasks_v1_tasksservice_createtask"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          post("/v1/tasks").payload(body)
        when "api_tasks_v1_tasksservice_deletetask"
          excludeKeys = ["id"]
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          delete("/v1/tasks/#{input['id']}").params(body)
        when "api_tasks_v1_tasksservice_gettask"
          excludeKeys = ["id"]
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          get("/v1/tasks/#{input['id']}").params(body)
        when "api_tasks_v1_tasksservice_updatetask"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          put("/v1/tasks/#{input['id']}").payload(body)
        end
      end,
      output_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_tasksservice_createtask"
          object_definitions["api.tasks.v1.CreateTaskResponse"]
        when "api_tasks_v1_tasksservice_deletetask"
          object_definitions["api.tasks.v1.DeleteTaskResponse"]
        when "api_tasks_v1_tasksservice_gettask"
          object_definitions["api.tasks.v1.GetTaskResponse"]
        when "api_tasks_v1_tasksservice_updatetask"
          object_definitions["api.tasks.v1.UpdateTaskResponse"]
        end
      end,
    },
  
    "action_task_comments": {
      title: "Task Comments",
      subtitle: "Interact with Task Comments in iAuditor",
      description: lambda do |input, picklist_label|
        "<span class='provider'>#{picklist_label['action_name'] || 'Interact with Task Comments'}</span> in <span class='provider'>iAuditor</span>"
      end,
      config_fields: [
        {
          name: "action_name",
          label: "Action",
          type: "string",
          
          control_type: "select",
          pick_list: "action_name_task_comments",
          sticky: true,
          
        },
      ],
      input_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_tasksservice_addcomment"
          object_definitions["api.tasks.v1.AddCommentRequest"]
        when "api_tasks_v1_tasksservice_updatecomment"
          object_definitions["api.tasks.v1.UpdateCommentRequest"]
        end
      end,
      execute: lambda do |connection, input, eis, eos, continue|
        case input['action_name']
        when "api_tasks_v1_tasksservice_addcomment"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          post("/v1/tasks/#{input['task_id']}/comment").payload(body)
        when "api_tasks_v1_tasksservice_updatecomment"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          put("/v1/tasks/#{input['task_id']}/comment/#{input['comment_id']}").payload(body)
        end
      end,
      output_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_tasksservice_addcomment"
          object_definitions["api.tasks.v1.AddCommentResponse"]
        when "api_tasks_v1_tasksservice_updatecomment"
          object_definitions["api.tasks.v1.UpdateCommentResponse"]
        end
      end,
    },
  
  },

  # Dynamic webhook example. Subscribes and unsubscribes webhooks programmatically
  # see more at https://docs.workato.com/developing-connectors/sdk/guides/building-triggers/dynamic-webhook.html
  triggers: {  
      trigger_inspections: {
         title: "Inspection Event",
  
         description: "<span class='provider'>Trigger for Inspection Event</span>",
  
         input_fields: lambda do |object_definitions|
          object_definitions["api.tasks.v1.TriggerInspectionRequest"]
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
          object_definitions["api.tasks.v1.TriggerInspectionResponse"]
         end,
      },
  },

  pick_lists: {  
    "dynamic_api_tasks_v1_tasksservice_listtasks": lambda do    
      body = {}
      resp = post("/api.tasks.v1.TasksService/ListTasks").payload(body)
      
      resp['tasks'].pluck('name', 'id')
    end,
    "action_name_tasks": lambda do
      [
        ["Create a new task", "api_tasks_v1_tasksservice_createtask"],
        ["Get a task by ID", "api_tasks_v1_tasksservice_gettask"],
        ["Update a task by ID", "api_tasks_v1_tasksservice_updatetask"],
        ["Delete a task by ID", "api_tasks_v1_tasksservice_deletetask"],
      ]
    end,
    "action_name_task_comments": lambda do
      [
        ["Add a comment to a task", "api_tasks_v1_tasksservice_addcomment"],
        ["Update a comment on a task", "api_tasks_v1_tasksservice_updatecomment"],
      ]
    end,
    "enum_api_tasks_v1_triggerinspectionrequest_triggerevent": lambda do
      [
        ["Inspection Started", "TRIGGER_EVENT_INSPECTION_STARTED"],
        ["Inspection Updated", "TRIGGER_EVENT_INSPECTION_UPDATED"],
        ["Inspection Completed", "TRIGGER_EVENT_INSPECTION_COMPLETED"],
        ["Inspection Archived", "TRIGGER_EVENT_INSPECTION_ARCHIVED"],
        ["Inspection Unarchived", "TRIGGER_EVENT_INSPECTION_UNARCHIVED"],
      ]
    end,
  },

  # Reusable methods can be called from object_definitions, picklists or actions
  # See more at https://docs.workato.com/developing-connectors/sdk/sdk-reference/methods.html
  methods: {
  }
}
