package genworkato

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"strings"

	tmpl "text/template"

	"github.com/Masterminds/sprig"
	"github.com/SafetyCulture/protoc-gen-workato/config"
	"github.com/SafetyCulture/protoc-gen-workato/template"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

func Escape(s string) string {
	return tmpl.HTMLEscapeString(s)
}

func EscapeActionName(s string) string {
	replacer := strings.NewReplacer(
		".", "_",
		"/", "_",
		" ", "_",
		"&", "and",
	)

	return replacer.Replace(s)
}

func formatStringSlice(slc []string) string {
	if slc == nil {
		return "[]"
	}
	b, err := json.Marshal(slc)
	if err != nil {
		return "[]"
	}
	return string(b)
}

var funcMap = tmpl.FuncMap{
	"escape":            Escape,
	"escapeActionName":  EscapeActionName,
	"formatStringSlice": formatStringSlice,
}

type Action struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

//go:embed templates/connector.rb.tmpl
var connectorTmpl string

func GenerateWorkatoConnector(gendoctemplate *gendoc.Template, cfg *config.Config) ([]byte, error) {
	workatoTemplate := template.FromGenDoc(gendoctemplate, cfg)

	tp, err := tmpl.New("Connector Template").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(connectorTmpl)
	if err != nil {
		return nil, err
	}

	// We should sort all of the lists in this template before rendering
	// This way the result is deterministic and diffing is easy.
	var buf bytes.Buffer
	err = tp.Execute(&buf, workatoTemplate)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
