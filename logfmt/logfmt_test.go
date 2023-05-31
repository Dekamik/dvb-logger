package logfmt

import (
    "bytes"
    dvblogger "dvb-logger/logger"
    "log"
    "strings"
    "testing"
)

func createLogger(buff *bytes.Buffer) *log.Logger {
    return log.New(buff, "", 0)
}

func TestLog(t *testing.T) {
    var buff bytes.Buffer

    cases := []struct {
        logger        dvblogger.Logger
        msgLevel      int64
        msgText       string
        msgProperties map[string]string
        expected      string
    }{
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), dvblogger.INFORMATION, "AnyMessage", nil, "level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), dvblogger.DEBUG, "AnyMessage", nil, ""},
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), dvblogger.WARNING, "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.WARNING, nil), dvblogger.INFORMATION, "AnyMessage", nil, ""},
        {New(createLogger(&buff), dvblogger.WARNING, nil), dvblogger.WARNING, "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.WARNING, nil), dvblogger.ERROR, "AnyMessage", nil, "level=error msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, map[string]string{"any_key": "AnyValue", "any key2": "Any value"}), dvblogger.INFORMATION, "AnyMessage", nil, "any_key2=\"Any value\" any_key=AnyValue level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), dvblogger.INFORMATION, "AnyMessage", map[string]string{"any_prop": "AnyValue", "any Prop": "Any thing"}, "any_Prop=\"Any thing\" any_prop=AnyValue level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, map[string]string{"global_prop": "AnyValue"}), dvblogger.INFORMATION, "AnyMessage", map[string]string{"msg_prop": "AnyValue"}, "global_prop=AnyValue msg_prop=AnyValue level=info msg=\"AnyMessage\""},
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
        logger   dvblogger.Logger
        msgText  string
        msgProps map[string]string
        expected string
    }{
        {New(createLogger(&buff), dvblogger.TRACE, nil), "AnyMessage", nil, "level=trace msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.TRACE, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=trace msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.TRACE, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=trace msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.TRACE, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=trace msg=\"AnyMessage\""},
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
        logger   dvblogger.Logger
        msgText  string
        msgProps map[string]string
        expected string
    }{
        {New(createLogger(&buff), dvblogger.DEBUG, nil), "AnyMessage", nil, "level=debug msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.DEBUG, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=debug msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.DEBUG, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=debug msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.DEBUG, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=debug msg=\"AnyMessage\""},
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
        logger   dvblogger.Logger
        msgText  string
        msgProps map[string]string
        expected string
    }{
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), "AnyMessage", nil, "level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=info msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.INFORMATION, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=info msg=\"AnyMessage\""},
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
        logger   dvblogger.Logger
        msgText  string
        msgProps map[string]string
        expected string
    }{
        {New(createLogger(&buff), dvblogger.WARNING, nil), "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.WARNING, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=warn msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.WARNING, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=warn msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.WARNING, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=warn msg=\"AnyMessage\""},
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
        logger   dvblogger.Logger
        msgText  string
        msgProps map[string]string
        expected string
    }{
        {New(createLogger(&buff), dvblogger.ERROR, nil), "AnyMessage", nil, "level=error msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", nil, "global_prop=value level=error msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.ERROR, nil), "AnyMessage", map[string]string{"msg_prop": "value"}, "msg_prop=value level=error msg=\"AnyMessage\""},
        {New(createLogger(&buff), dvblogger.ERROR, map[string]string{"global_prop": "value"}), "AnyMessage", map[string]string{"msg_prop": "value"}, "global_prop=value msg_prop=value level=error msg=\"AnyMessage\""},
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
