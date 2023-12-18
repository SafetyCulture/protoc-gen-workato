package template

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SafetyCulture/protoc-gen-workato/template/schema"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	extensions "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http"
)

const pbRepeated = "repeated"

const errorHandling = `.after_error_response(/.*/) do |code, body, headers, message| call('after_error_response', code, body, headers, message) end`

// Used to identify parameters in a path e.g. `/users/{used_id}`
var paramMatch = regexp.MustCompile(`({[\.\w]+})`)

func hasRepeatedType(message *gendoc.Message) bool {
	if !message.HasFields {
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

					vals := strings.Split(param, ".")
					for i, val := range vals {
						vals[i] = fmt.Sprintf(":%s", val) // Convert the string into a ruby symbol
					}

					path = strings.ReplaceAll(path, match, fmt.Sprintf("#{input.dig(%s)}", strings.Join(vals, ", ")))
				}
			}

			if rule.Body == "*" {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Func:             fmt.Sprintf(`%s("%s").payload(body)%s`, mthd, path, errorHandling),
				}
			}

			if rule.Body != "" {
				return schema.ExecCode{
					ExcludeFromQuery: append(params, rule.Body),
					Func:             fmt.Sprintf(`%s("%s").payload(input['%s']).params(body)%s`, mthd, path, rule.Body, errorHandling),
				}
			}

			if hasRepeatedType(t.messageMap[method.RequestFullType]) {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Body:             "qparams = call('encode_array_to_query_params', body)",
					Func:             fmt.Sprintf(`%s("%s?#{qparams}")%s`, mthd, path, errorHandling),
				}
			} else {
				return schema.ExecCode{
					ExcludeFromQuery: params,
					Func:             fmt.Sprintf(`%s("%s").params(body)%s`, mthd, path, errorHandling),
				}
			}
		}
	}

	return schema.ExecCode{
		Func: fmt.Sprintf(`post("/%s/%s").payload(body)%s`, service.FullName, method.Name, errorHandling),
	}
}
