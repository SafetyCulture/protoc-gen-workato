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
            optional: "false",
            type: "string",
            hint: "The name of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
        ]
      end
    },
  
    "api.tasks.v1.GetTaskResponse": {
      fields: lambda do |connection, config_fields, object_definitions|
        [
          {
            name: "id",
            label: "Id",
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
          {
            name: "name",
            label: "Name",
            optional: "false",
            type: "string",
            hint: "The name of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
          {
            name: "name",
            label: "Name",
            optional: "false",
            type: "string",
            hint: "The name of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
          {
            name: "comment",
            label: "Comment",
            optional: "false",
            type: "string",
            hint: "The comment to add",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the comment",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
            optional: "false",
            type: "string",
            hint: "The ID of the task",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
          {
            name: "comment_id",
            label: "Comment Id",
            optional: "false",
            type: "string",
            hint: "The ID comment to update",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
          },
          {
            name: "comment",
            label: "Comment",
            optional: "false",
            type: "string",
            hint: "The updated comment",
            
            
            
            #toggle_hint: "",
            #toggle_field: "<nil>",
            
            pick_list: "",
            #delimiter: "",
            sticky: true,
            #render_input: "",
            #parse_output: "",
            #change_on_blur: "false",
            #support_pills: "false",
            #custom: "false",
            #extends_schema: "false",
            #list_mode: "",
            #list_mode_toggle: "false",
            #item_label: "",
            #add_field_label: "",
            #empty_schema_message: "",
            #sample_data_type: "",
            #ng_if: "",
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
  
  },

  actions: {  
    "Tasks": {
      title: "Tasks",
      subtitle: "Interact with Tasks in iAuditor",
      description: lambda do |input, picklist_label|
        "<span class='provider'>#{picklist_label['action_name'] || 'Interact with Tasks'}</span> in <span class='provider'>iAuditor</span>"
      end,
      config_fields: [
        {
          name: "action_name",
          label: "Action",
          optional: "false",
          type: "string",
          hint: "",
          
          
          control_type: "select",
          #toggle_hint: "",
          #toggle_field: "<nil>",
          
          pick_list: "action_name_Tasks",
          #delimiter: "",
          sticky: true,
          #render_input: "",
          #parse_output: "",
          #change_on_blur: "false",
          #support_pills: "false",
          #custom: "false",
          #extends_schema: "false",
          #list_mode: "",
          #list_mode_toggle: "false",
          #item_label: "",
          #add_field_label: "",
          #empty_schema_message: "",
          #sample_data_type: "",
          #ng_if: "",
        },
      ],
      input_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_TasksService_CreateTask"
          object_definitions["api.tasks.v1.CreateTaskRequest"]
        when "api_tasks_v1_TasksService_DeleteTask"
          object_definitions["api.tasks.v1.DeleteTaskRequest"]
        when "api_tasks_v1_TasksService_GetTask"
          object_definitions["api.tasks.v1.GetTaskRequest"]
        when "api_tasks_v1_TasksService_UpdateTask"
          object_definitions["api.tasks.v1.UpdateTaskRequest"]
        end
      end,
      execute: lambda do |connection, input, eis, eos, continue|
        case input['action_name']
        when "api_tasks_v1_TasksService_CreateTask"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          post("/v1/tasks").payload(body)
        when "api_tasks_v1_TasksService_DeleteTask"
          excludeKeys = ["id"]
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          delete("/v1/tasks/#{input['id']}").params(body)
        when "api_tasks_v1_TasksService_GetTask"
          excludeKeys = ["id"]
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          get("/v1/tasks/#{input['id']}").params(body)
        when "api_tasks_v1_TasksService_UpdateTask"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          put("/v1/tasks/#{input['id']}").payload(body)
        end
      end,
      output_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_TasksService_CreateTask"
          object_definitions["api.tasks.v1.CreateTaskResponse"]
        when "api_tasks_v1_TasksService_DeleteTask"
          object_definitions["api.tasks.v1.DeleteTaskResponse"]
        when "api_tasks_v1_TasksService_GetTask"
          object_definitions["api.tasks.v1.GetTaskResponse"]
        when "api_tasks_v1_TasksService_UpdateTask"
          object_definitions["api.tasks.v1.UpdateTaskResponse"]
        end
      end,
    },
  
    "Task_Comments": {
      title: "Task Comments",
      subtitle: "Interact with Task Comments in iAuditor",
      description: lambda do |input, picklist_label|
        "<span class='provider'>#{picklist_label['action_name'] || 'Interact with Task Comments'}</span> in <span class='provider'>iAuditor</span>"
      end,
      config_fields: [
        {
          name: "action_name",
          label: "Action",
          optional: "false",
          type: "string",
          hint: "",
          
          
          control_type: "select",
          #toggle_hint: "",
          #toggle_field: "<nil>",
          
          pick_list: "action_name_Task_Comments",
          #delimiter: "",
          sticky: true,
          #render_input: "",
          #parse_output: "",
          #change_on_blur: "false",
          #support_pills: "false",
          #custom: "false",
          #extends_schema: "false",
          #list_mode: "",
          #list_mode_toggle: "false",
          #item_label: "",
          #add_field_label: "",
          #empty_schema_message: "",
          #sample_data_type: "",
          #ng_if: "",
        },
      ],
      input_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_TasksService_AddComment"
          object_definitions["api.tasks.v1.AddCommentRequest"]
        when "api_tasks_v1_TasksService_UpdateComment"
          object_definitions["api.tasks.v1.UpdateCommentRequest"]
        end
      end,
      execute: lambda do |connection, input, eis, eos, continue|
        case input['action_name']
        when "api_tasks_v1_TasksService_AddComment"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          post("/v1/tasks/#{input['task_id']}/comment").payload(body)
        when "api_tasks_v1_TasksService_UpdateComment"
          excludeKeys = []
          body = input.select {|k,v| k != "action_name" and not excludeKeys.include? k }
          put("/v1/tasks/#{input['task_id']}/comment/#{input['comment_id']}").payload(body)
        end
      end,
      output_fields: lambda do |object_definitions, connection, config_fields|
        case config_fields['action_name']
        when "api_tasks_v1_TasksService_AddComment"
          object_definitions["api.tasks.v1.AddCommentResponse"]
        when "api_tasks_v1_TasksService_UpdateComment"
          object_definitions["api.tasks.v1.UpdateCommentResponse"]
        end
      end,
    },
  
  },

  triggers: {
  },

  pick_lists: {  
  "action_name_Tasks": lambda do
    [
      ["Create a new task", "api_tasks_v1_TasksService_CreateTask"],
      ["Get a task by ID", "api_tasks_v1_TasksService_GetTask"],
      ["Update a task by ID", "api_tasks_v1_TasksService_UpdateTask"],
      ["Delete a task by ID", "api_tasks_v1_TasksService_DeleteTask"],
    ]
  end,
  "action_name_Task_Comments": lambda do
    [
      ["Add a comment to a task", "api_tasks_v1_TasksService_AddComment"],
      ["Update a comment on a task", "api_tasks_v1_TasksService_UpdateComment"],
    ]
  end,
  },

  # Reusable methods can be called from object_definitions, picklists or actions
  # See more at https://docs.workato.com/developing-connectors/sdk/sdk-reference/methods.html
  methods: {
  }
}
