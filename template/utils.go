package template

import (
	"fmt"
	"strings"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

const unspecified = "_UNSPECIFIED" //add _UNKNOWN

func escapeKeyName(s string) string {
	replacer := strings.NewReplacer(
		".", "_",
		"/", "_",
		" ", "_",
		"&", "and",
	)

	return strings.ToLower(replacer.Replace(s))
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

// shouldIncludeEnum returns true if the name doesn't end in _UNSPECIFIED.
// otherwise, returns false
func shouldIncludeEnum(enum *gendoc.EnumValue) bool {
	return !strings.HasSuffix(enum.Name, unspecified)
}

// removeUnspecifiedValue removes the _UNSPECIFIED if at index 0
func removeUnspecifiedValue(enum *gendoc.Enum) {
	if len(enum.Values) >= 1 && !shouldIncludeEnum(enum.Values[0]) {
		enum.Values = enum.Values[1:]
	}
}
