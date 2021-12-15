package genworkato

import (
	"os"

	"github.com/SafetyCulture/protoc-gen-workato/config"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protokit"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"gopkg.in/yaml.v3"
)

// PluginOptions encapsulates options for the plugin. The type of renderer, template file, and the name of the output
// file are included.
type PluginOptions struct {
	ConfigFile string
}

// SupportedFeatures describes a flag setting for supported features.
var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// Plugin describes a protoc code generate plugin. It's an implementation of Plugin from github.com/pseudomuto/protokit
type Plugin struct{}

// Generate compiles the documentation and generates the CodeGeneratorResponse to send back to protoc. It does this
// by rendering a template based on the options parsed from the CodeGeneratorRequest.
func (p *Plugin) Generate(r *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	options, err := ParseOptions(r)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(options.ConfigFile)
	if err != nil {
		return nil, err
	}

	var configYaml config.Config
	err = yaml.NewDecoder(f).Decode(&configYaml)
	if err != nil {
		return nil, err
	}

	result := protokit.ParseCodeGenRequest(r)

	resp := new(pluginpb.CodeGeneratorResponse)
	template := gendoc.NewTemplate(result)

	workatoOutput, err := GenerateWorkatoConnector(template, &configYaml)
	if err != nil {
		return nil, err
	}

	resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String("workato_connector.rb"),
		Content: proto.String(string(workatoOutput)),
	})

	resp.SupportedFeatures = proto.Uint64(SupportedFeatures)

	return resp, nil
}

// ParseOptions parses plugin options from a CodeGeneratorRequest. It does this by splitting the `Parameter` field from
// the request object and parsing out the type of renderer to use and the name of the file to be generated.
//
// The parameter (`--workato_opt`) must be of the format <config_file_path>.
// The file will be written to the directory specified with the `--workato_out` argument to protoc.
func ParseOptions(req *pluginpb.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		ConfigFile: "config.yaml",
	}

	params := req.GetParameter()
	if params == "" {
		return options, nil
	}

	options.ConfigFile = params

	return options, nil
}
