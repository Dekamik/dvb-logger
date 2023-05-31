package logfmt

import (
    "testing"
)

func TestMapPropertiesToLogFmt(t *testing.T) {
    properties := map[string]string{
        "a": "one",
        "b": "two",
        "c": "three",
    }
    expected := "a=one b=two c=three "
    actual := mapPropertiesToLogFmt(properties)

    if actual != expected {
        t.Errorf("mapPropertiesToLogFmt(%v) should return %s but got %s", properties, expected, actual)
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
            t.Errorf("addQuotes(%s) should return %s, but got %s", c.from, c.expected, str)
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
            t.Errorf("replaceSpaces(%s) should return %s, but got %s", c.from, c.expected, str)
        }
    }
}
