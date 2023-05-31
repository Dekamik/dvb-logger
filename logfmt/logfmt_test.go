package logfmt

import (
	"bytes"
	dvb_logger "dvb-logger"
	"log"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	var (
		buff       bytes.Buffer
		mockLogger = log.New(&buff, "", 0)
	)

	cases := []struct {
		logger        dvb_logger.Logger
		msgLevel      int64
		msgText       string
		msgProperties *map[string]string
		expected      string
	}{
		{New(mockLogger, dvb_logger.INFORMATION, nil), dvb_logger.INFORMATION, "AnyMessage", nil, "level=info msg=\"AnyMessage\""},
		{New(mockLogger, dvb_logger.INFORMATION, nil), dvb_logger.DEBUG, "AnyMessage", nil, ""},
		{New(mockLogger, dvb_logger.INFORMATION, nil), dvb_logger.WARNING, "AnyMessage", nil, "level=warn msg=\"AnyMessage\""},
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
