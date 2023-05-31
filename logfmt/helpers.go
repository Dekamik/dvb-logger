package logfmt

import (
	"fmt"
	"strings"
)

func mapPropertiesToLogFmt(properties map[string]string) string {
	var s = ""
	for k, v := range properties {
		s += fmt.Sprintf("%s=%s ", replaceSpaces(k), addQuotes(v))
	}
	return s
}

func addQuotes(s string) string {
	if strings.Contains(s, " ") {
		return fmt.Sprintf("\"%s\"", s)
	}
	return s
}

func replaceSpaces(s string) string {
	if strings.Contains(s, " ") {
		return strings.Replace(s, " ", "_", -1)
	}
	return s
}
