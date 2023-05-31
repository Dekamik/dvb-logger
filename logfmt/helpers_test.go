package logfmt

import (
	"strings"
	"testing"
)

func TestMapPropertiesToLogFmt(t *testing.T) {
	properties := map[string]string{
		"a": "one",
		"b": "two",
		"c": "three",
	}
	actual := mapPropertiesToLogFmt(properties)

	if !strings.Contains(actual, "a=one") {
		t.Errorf("expected \"%s\" in \"%s\"", "a=one", actual)
	}
	if !strings.Contains(actual, "b=two") {
		t.Errorf("expected \"%s\" in \"%s\"", "b=two", actual)
	}
	if !strings.Contains(actual, "c=three") {
		t.Errorf("expected \"%s\" in \"%s\"", "c=three", actual)
	}
}

func TestAddQuotes(t *testing.T) {
	cases := []struct{ from, expected string }{
		{"any", "any"},
		{"any string", "\"any string\""},
	}

	for _, c := range cases {
		str := addQuotes(c.from)
		if str != c.expected {
			t.Errorf("addQuotes(%s) should be %s, but got %s", c.from, c.expected, str)
		}
	}
}

func TestReplaceSpaces(t *testing.T) {
	cases := []struct{ from, expected string }{
		{"any", "any"},
		{"any string", "any_string"},
	}

	for _, c := range cases {
		str := replaceSpaces(c.from)
		if str != c.expected {
			t.Errorf("replaceSpaces(%s) should be %s, but got %s", c.from, c.expected, str)
		}
	}
}
