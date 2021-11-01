package template

import (
	"fmt"
	"strings"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

func escapeKeyName(s string) string {
	replacer := strings.NewReplacer(
		".", "_",
		"/", "_",
		" ", "_",
		"&", "and",
	)

	return replacer.Replace(s)
}

func enumPicklistName(enum *gendoc.Enum) string {
	return fmt.Sprintf("%s_%s", "enum", escapeKeyName(enum.FullName))
}

func actionPicklistName(group string) string {
	return fmt.Sprintf("%s_%s", "action_name", escapeKeyName(group))
}

func fieldTitleFromName(name string) string {
	return strings.Title(strings.ReplaceAll(name, "_", " "))
}
