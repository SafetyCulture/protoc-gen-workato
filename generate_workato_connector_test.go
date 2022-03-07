package genworkato_test

import (
	"strings"
	"testing"

	genworkato "github.com/SafetyCulture/protoc-gen-workato"
	"github.com/SafetyCulture/protoc-gen-workato/config"
	"github.com/SafetyCulture/protoc-gen-workato/template/schema"

	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/google_api"           // imported for side effects
	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_openapiv2" // imported for side effects
	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_workato"   // imported for side effects
	"github.com/bradleyjkemp/cupaloy"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	_ "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http" // imported for side effects
	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	"gotest.tools/assert"
)

//go:generate buf build -o fixtures/image.bin

func TestGenerateWorkatoConnector(t *testing.T) {
	set, err := utils.LoadDescriptorSet("fixtures", "image.bin")
	assert.NilError(t, err)

	req := utils.CreateGenRequest(set, "fixtures/tasks.proto", "google/protobuf/empty.proto")
	result := protokit.ParseCodeGenRequest(req)

	template := gendoc.NewTemplate(result)

	content, err := genworkato.GenerateWorkatoConnector(template, &config.Config{
		Name:             "My Workato Connector",
		AppBaseURL:       "https://api.example.com",
		DeveloperDocsURL: "https://developer.example.com",
		TemplateFile:     "fixtures/connector_template.rb",
		Action: map[string]config.Action{
			"Tasks": {
				DefaultHelpMessage: &schema.HelpMessage{
					Body:          "Tasks allow you to define work to be done an assign it to someone.",
					LearnMoreText: "Learn more about tasks",
					LearnMoreURL:  "https://google.com",
				},
				InputFields: []schema.FieldDefinition{
					{
						Name: "custom_field",
						Type: "text",
						NgIf: "input['action'] == 'api_tasks_v1_tasksservice_customaction'",
					},
				},
			},
		},
		Method: map[string]config.Method{
			"api.tasks.v1.TasksService/CustomAction": {
				Exec: "# does a thing",
			},
		},
		Message: map[string]config.Message{
			"api.tasks.v1.CustomActionRequest": {
				Exec: `data = get("/data/for_tasks/#{input['custom_field']}")
data.map ...`,
			},
			"api.tasks.v1.UnusedButIncluded": {
				Include: true,
			},
		},
		CustomMethods: []*schema.MethodDefinition{
			{
				Name:   "does_a_thing",
				Params: []string{"param_1", "param_2"},
				Exec:   "param_1 + param_2",
			},
			{
				Name:   "does_another_thing",
				Params: []string{"param_1", "param_2"},
				Exec: `get("/an/api/#{param_1}")
.body(param_2)`,
			},
		},
		VisibilityRestrictionSelectors: []string{
			"WORKATO",
		},
	})
	assert.NilError(t, err)

	cupaloy.SnapshotT(t, strings.TrimSpace(string(content)))
}
