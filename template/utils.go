package template

import (
	"fmt"
	"strings"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

const (
	unspecified = "_UNSPECIFIED"
	unknown     = "_UNKNOWN"
)

func escapeKeyName(s string) string {
	replacer := strings.NewReplacer(
		".", "_",
		"/", "_",
		" ", "_",
		"&", "and",
	)

	return strings.ToLower(replacer.Replace(s))
}

func fullActionName(service *gendoc.Service, method *gendoc.ServiceMethod) string {
	return escapeKeyName(fmt.Sprintf("%s/%s", service.FullName, method.Name))
}

func enumPicklistName(enum *gendoc.Enum) string {
	return fmt.Sprintf("%s_%s", "enum", escapeKeyName(enum.FullName))
}

func actionPicklistName(group string) string {
	return fmt.Sprintf("%s_%s", "action_name", escapeKeyName(group))
}

func dynamicPicklistName(actionName string) string {
	return fmt.Sprintf("%s_%s", "dynamic", actionName)
}

func fieldTitleFromName(name string) string {
	return strings.Title(strings.ReplaceAll(name, "_", " "))
}

// enumValueShouldBeExcluded returns true if the name ends in _UNSPECIFIED or in _UNKNOWN.
// otherwise, returns false
func enumValueShouldBeExcluded(enum *gendoc.EnumValue) bool {
	if strings.HasSuffix(enum.Name, unspecified) || strings.HasSuffix(enum.Name, unknown) {
		return true
	}
	return false
}

// removeUnspecifiedValue removes the _UNSPECIFIED or _UNKNOWN if exists at index 0
func removeUnspecifiedValue(enum *gendoc.Enum) {
	if len(enum.Values) >= 1 && enumValueShouldBeExcluded(enum.Values[0]) {
		enum.Values = enum.Values[1:]
	}
}
