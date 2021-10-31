package genworkato

import (
	"bytes"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	tmpl "text/template"

	"github.com/Masterminds/sprig"
	"github.com/SafetyCulture/protoc-gen-workato/config"
	"github.com/SafetyCulture/protoc-gen-workato/template"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

func formatStringSlice(slc []string) (string, error) {
	if slc == nil {
		return "[]", nil
	}
	b, err := json.Marshal(slc)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Action struct {
	Service *gendoc.Service
	Method  *gendoc.ServiceMethod
}

//go:embed templates/*.tmpl.rb
var tmpls embed.FS

// GenerateWorkatoConnector generates a Workato SDK Connector from protobufs
func GenerateWorkatoConnector(gendoctemplate *gendoc.Template, cfg *config.Config) ([]byte, error) {
	workatoTemplate := template.FromGenDoc(gendoctemplate, cfg)

	tp := tmpl.New("Connector Template").
		Funcs(sprig.TxtFuncMap())

	tp.Funcs(tmpl.FuncMap{
		"format_string_slice": formatStringSlice,
		"field_key": func(name string, data interface{}) (string, error) {
			value := ""
			switch d := data.(type) {
			case string:
				if len(d) != 0 {
					value = fmt.Sprintf("\"%s\"", strings.ReplaceAll(d, `"`, `\"`))
				}
			}

			if len(value) == 0 {
				return "", nil
			}

			return fmt.Sprintf("%s: %s,\n  ", name, value), nil
		},
		"include": func(name string, data interface{}) (string, error) {
			buf := bytes.NewBuffer(nil)
			if err := tp.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	})

	var err error
	if tp, err = tp.ParseFS(tmpls, "templates/*.tmpl.rb"); err != nil {
		return nil, err
	}

	// We should sort all of the lists in this template before rendering
	// This way the result is deterministic and diffing is easy.
	var buf bytes.Buffer
	err = tp.ExecuteTemplate(&buf, "connector.tmpl.rb", workatoTemplate)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
