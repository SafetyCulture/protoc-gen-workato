package template

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	extensions "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http"
)

const (
	httpGet    = "get"
	pbRepeated = "repeated"
)

// Used to identify parameters in a path e.g. `/users/{used_id}`
var paramMatch = regexp.MustCompile(`({\w+})`)

func hasRepeatedType(message *gendoc.Message) bool {
	if message.HasFields == false {
		return false
	}

	for _, s := range message.Fields {
		if s.Label == pbRepeated {
			return true
		}
	}
	return false
}

func (t *WorkatoTemplate) getExecuteCode(service *gendoc.Service, method *gendoc.ServiceMethod) schema.ExecCode {
	if override, ok := t.config.Method[fmt.Sprintf("%s/%s", service.FullName, method.Name)]; ok {
		return schema.ExecCode{
			Func: override.Exec,
		}
	}

	if opts, ok := method.Option("google.api.http").(extensions.HTTPExtension); ok {
		if len(opts.Rules) != 0 {
			rule := opts.Rules[0]

			mthd := strings.ToLower(rule.Method)
			path := rule.Pattern

			params := []string{}
			if matches := paramMatch.FindAllString(path, -1); len(matches) != 0 {
				for _, match := range matches {
					param := match[1 : len(match)-1]
					params = append(params, param)

					path = strings.ReplaceAll(path, match, fmt.Sprintf("#{input['%s']}", param))
				}
			}

			if rule.Body == "*" {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Func:             fmt.Sprintf(`%s("%s").payload(body)`, mthd, path),
				}
			}

			if rule.Body != "" {
				return schema.ExecCode{
					ExcludeFromQuery: append(params, rule.Body),
					Func:             fmt.Sprintf(`%s("%s").payload(input['%s']).params(body)`, mthd, path, rule.Body),
				}
			}

			if mthd == httpGet && hasRepeatedType(t.messageMap[method.RequestFullType]) {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Body:             "qparams = call('encode_array_to_query_params', body)",
					Func:             fmt.Sprintf(`%s("%s?#{qparams}")`, mthd, path),
				}
			} else {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Func:             fmt.Sprintf(`%s("%s").params(body)`, mthd, path),
				}
			}
		}
	}

	return schema.ExecCode{
		Func: fmt.Sprintf(`post("/%s/%s").payload(body)`, service.FullName, method.Name),
	}
}
