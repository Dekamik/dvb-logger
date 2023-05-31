package logfmt

import (
    "fmt"
    "sort"
    "strings"
)

// mapPropertiesToLogFmt converts a map[string]string to logfmt like "key=value key=value"
func mapPropertiesToLogFmt(properties map[string]string) string {
    var (
        s    = ""
        keys = make([]string, len(properties))
    )

    // This is done to enforce deterministic behaviour on string conversion, since maps are randomly ordered and may
    // not access objects in the order which the user specified them on map initialization.
    i := 0
    for k := range properties {
        keys[i] = k
        i++
    }
    sort.Strings(keys)

    for _, k := range keys {
        s += fmt.Sprintf("%s=%s ", replaceSpaces(k), addQuotes(properties[k]))
    }

    return s
}

// addQuotes surrounds the string with quotation marks if any space is found in the string.
// If no spaces are found, the string is returned as-is.
func addQuotes(s string) string {
    if strings.Contains(s, " ") {
        return fmt.Sprintf("\"%s\"", s)
    }
    return s
}

// replaceSpaces replaces all whitespace in the string with an underscore.
// If no spaces are found, the string is returned as-is.
func replaceSpaces(s string) string {
    if strings.Contains(s, " ") {
        return strings.Replace(s, " ", "_", -1)
    }
    return s
}
