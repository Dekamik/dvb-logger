package logfmt

import (
	"bytes"
	"dvb-logger/logger"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func createLogger(buff *bytes.Buffer) *log.Logger {
	return log.New(buff, "", 0)
}

func TestLog(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger        logger.Logger
		msgLevel      int64
		msgText       string
		msgProperties map[string]string
		expected      string
	}{
		{New(createLogger(&buff), logger.INFORMATION, nil), logger.INFORMATION, "AnyMessage", nil, "level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, nil), logger.DEBUG, "AnyMessage", nil, ""},
		{New(createLogger(&buff), logger.INFORMATION, nil), logger.WARNING, "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.WARNING, nil), logger.INFORMATION, "AnyMessage", nil, ""},
		{New(createLogger(&buff), logger.WARNING, nil), logger.WARNING, "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.WARNING, nil), logger.ERROR, "AnyMessage", nil, "level=error msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, map[string]string{"any_key": "AnyValue", "any key2": "Any value"}), logger.INFORMATION, "AnyMessage", nil, "any_key2=\"Any value\" any_key=AnyValue level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, nil), logger.INFORMATION, "AnyMessage", map[string]string{"any_prop": "AnyValue", "any Prop": "Any thing"}, "any_Prop=\"Any thing\" any_prop=AnyValue level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, map[string]string{"global_prop": "AnyValue"}), logger.INFORMATION, "AnyMessage", map[string]string{"msg_prop": "AnyValue"}, "global_prop=AnyValue msg_prop=AnyValue level=info msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.Log(c.msgLevel, c.msgText, c.msgProperties)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("Log(%d, %s, %v) should write \"%s\", but got \"%s\"", c.msgLevel, c.msgText, c.msgProperties, c.expected, buff.String())
		}
	}
}

func TestLogTrace(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.TRACE, nil), "AnyMessage", nil, "level=trace msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.TRACE, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=trace msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.TRACE, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=trace msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.TRACE, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=trace msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.LogTrace(c.msgText, c.msgProps)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("LogTrace(%s, %v) should write \"%s\", but got \"%s\"", c.msgText, c.msgProps, c.expected, buff.String())
		}
	}
}

func TestLogDebug(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.DEBUG, nil), "AnyMessage", nil, "level=debug msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.DEBUG, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=debug msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.DEBUG, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=debug msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.DEBUG, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=debug msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.LogDebug(c.msgText, c.msgProps)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("LogDebug(%s, %v) should write \"%s\", but got \"%s\"", c.msgText, c.msgProps, c.expected, buff.String())
		}
	}
}

func TestLogInformation(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.INFORMATION, nil), "AnyMessage", nil, "level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=info msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.INFORMATION, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=info msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.LogInformation(c.msgText, c.msgProps)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("LogInformation(%s, %v) should write \"%s\", but got \"%s\"", c.msgText, c.msgProps, c.expected, buff.String())
		}
	}
}

func TestLogWarning(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.WARNING, nil), "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.WARNING, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=warn msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.WARNING, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=warn msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.WARNING, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=warn msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.LogWarning(c.msgText, c.msgProps)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("LogWarning(%s, %v) should write \"%s\", but got \"%s\"", c.msgText, c.msgProps, c.expected, buff.String())
		}
	}
}

func TestLogError(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", nil, "level=error msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=error msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=error msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=error msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		buff.Reset()
		c.logger.LogError(c.msgText, c.msgProps)
		lines := strings.Split(buff.String(), "\n")
		if lines[0] != c.expected {
			t.Errorf("LogError(%s, %v) should write \"%s\", but got \"%s\"", c.msgText, c.msgProps, c.expected, buff.String())
		}
	}
}

func TestLogFatal(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", nil, "level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=fatal msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		if os.Getenv("BE_FATAL") == "1" {
			buff.Reset()
			c.logger.LogFatal(c.msgText, c.msgProps)
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestLogFatal")
		cmd.Env = append(os.Environ(), "BE_FATAL=1")
		err := cmd.Run()

		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			return
		}
		t.Errorf("process ran with err %v, want exit status 1", err)
	}
}

func TestLogPanic(t *testing.T) {
	var buff bytes.Buffer

	cases := []struct {
		logger   logger.Logger
		msgText  string
		msgProps map[string]string
		expected string
	}{
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", nil, "level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=fatal msg=\"AnyMessage\""},
		{New(createLogger(&buff), logger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=fatal msg=\"AnyMessage\""},
	}

	for _, c := range cases {
		if os.Getenv("BE_PANIC") == "1" {
			buff.Reset()
			c.logger.LogPanic(c.msgText, c.msgProps)
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestLogPanic")
		cmd.Env = append(os.Environ(), "BE_PANIC=1")
		err := cmd.Run()

		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			return
		}
		t.Errorf("process ran with err %v, want exit status 1", err)
	}
}
