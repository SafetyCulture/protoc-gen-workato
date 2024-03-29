package genworkato

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	tmpl "text/template"

	"github.com/Masterminds/sprig/v3"
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

//go:embed templates/*.tmpl.rb
var tmpls embed.FS

// GenerateWorkatoConnector generates a Workato SDK Connector from protobufs
func GenerateWorkatoConnector(gendoctemplate *gendoc.Template, cfg *config.Config) ([]byte, error) {
	workatoTemplate, err := template.FromGenDoc(gendoctemplate, cfg)
	if err != nil {
		return nil, err
	}

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
			case *bool:
				if d != nil {
					value = "false"
					if *d {
						value = "true"
					}
				}
			case bool:
				value = "false"
				if d {
					value = "true"
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

	if tp, err = tp.ParseFS(tmpls, "templates/*.tmpl.rb"); err != nil {
		return nil, err
	}

	templateName := "connector.tmpl.rb"
	if cfg.TemplateFile != "" {
		templateName = "custom_template"
		customTp := tp.New(templateName)
		b, err := os.ReadFile(cfg.TemplateFile)
		if err != nil {
			return nil, err
		}
		_, err = customTp.Parse(string(b))
		if err != nil {
			return nil, err
		}
	}

	// We should sort all the lists in this template before rendering
	// This way the result is deterministic and diffing is easy.
	var buf bytes.Buffer
	err = tp.ExecuteTemplate(&buf, templateName, workatoTemplate)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
