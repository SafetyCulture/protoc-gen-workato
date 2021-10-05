package genworkato

import (
	"fmt"
	"regexp"
	"strings"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
	httpext "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http"
)

// Used to identify parameters in a path e.g. `/users/{used_id}`
var paramMatch = regexp.MustCompile(`({\w+})`)

func getExecuteCode(messages map[string]*gendoc.Message, service *gendoc.Service, method *gendoc.ServiceMethod) Endpoint {
	if opts, ok := method.Option("google.api.http").(httpext.HTTPExtension); ok {
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
				return Endpoint{
					Func: fmt.Sprintf(`%s("%s", body)`, mthd, path),
				}
			}

			if rule.Body != "" {
				return Endpoint{
					ExcludeFromQuery: append(params, rule.Body),
					Func:             fmt.Sprintf(`%s("%s?#{body.encode_www_form}", input['%s'])`, mthd, path, rule.Body),
				}
			}

			return Endpoint{
				ExcludeFromQuery: params,
				Func:             fmt.Sprintf(`%s("%s?#{body.encode_www_form}")`, mthd, path),
			}
		}
	}

	return Endpoint{
		Func: fmt.Sprintf(`post("/%s/%s", body)`, service.FullName, method.Name),
	}
}