package genworkato_test

import (
	"bytes"
	"os"
	"testing"

	genworkato "github.com/SafetyCulture/protoc-gen-workato"
	"github.com/SafetyCulture/protoc-gen-workato/config"

	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_openapiv2" // imported for side effects
	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_workato"   // imported for side effects
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

	req := utils.CreateGenRequest(set, "fixtures/tasks.proto")
	result := protokit.ParseCodeGenRequest(req)

	template := gendoc.NewTemplate(result)

	content, err := genworkato.GenerateWorkatoConnector(template, &config.Config{})
	assert.NilError(t, err)

	f, err := os.ReadFile("./fixtures/_generated/connector.rb")
	assert.NilError(t, err)

	var buf bytes.Buffer
	buf.Write(f)

	assert.Equal(t, buf.String(), string(content))
}
