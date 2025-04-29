package template

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gomarkdown/markdown"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

const (
	unspecified = "_UNSPECIFIED"
	unknown     = "_UNKNOWN"
)

var keyNameReplacer = strings.NewReplacer(
	".", "_",
	"/", "_",
	" ", "_",
	"&", "and",
	"(", "_",
	")", "_",
	"[", "_",
	"]", "_",
	"-", "_",
)

var stringValueReplacer = strings.NewReplacer(
	"'", `\'`,
	`"`, `\"`,
)

func escapeStringValue(s string) string {
	return stringValueReplacer.Replace(s)
}

func escapeKeyName(s string) string {
	return strings.ToLower(keyNameReplacer.Replace(s))
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

func getFieldTitle(field *gendoc.MessageField) string {
	c := cases.Title(language.AmericanEnglish, cases.NoLower)
	title := c.String(strings.ReplaceAll(field.Name, "_", " "))
	if field.Options["deprecated"] == true {
		title = title + "  ⛔ Deprecated - Please do not use ⛔ "
	}
	return title
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

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func markdownToHTML(md string) string {
	mdParser := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.HardLineBreak)
	renderer := html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	})

	if strings.TrimSpace(md) == "" {
		return ""
	}
	html := markdown.ToHTML([]byte(md), mdParser, renderer)
	return strings.TrimSpace(string(html))
}

func boolPtr(v bool) *bool {
	return &v
}
