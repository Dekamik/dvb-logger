package logfmt

import (
	"dvb-logger"
	"log"
)

const logFmtMessageFormat = "level=%s msg=\"%s\"\n"
const logFmtMessageFormatWithProps = "%slevel=%s msg=\"%s\"\n"

var _ dvb_logger.Logger = &Logger{}

type Logger struct {
	level  int64
	logger *log.Logger
}

func (l *Logger) Log(level int64, message string, properties *map[string]string) {
	if level >= l.level {
		if properties == nil {
			l.logger.Printf(logFmtMessageFormat, dvb_logger.LevelNames[level], message)
		} else {
			l.logger.Printf(logFmtMessageFormatWithProps, mapPropertiesToLogFmt(*properties), dvb_logger.LevelNames[level], message)
		}
	}
}

func (l *Logger) LogTrace(message string, properties *map[string]string) {
	l.Log(dvb_logger.TRACE, message, properties)
}

func (l *Logger) LogDebug(message string, properties *map[string]string) {
	l.Log(dvb_logger.DEBUG, message, properties)
}

func (l *Logger) LogInformation(message string, properties *map[string]string) {
	l.Log(dvb_logger.INFORMATION, message, properties)
}

func (l *Logger) LogWarning(message string, properties *map[string]string) {
	l.Log(dvb_logger.WARNING, message, properties)
}

func (l *Logger) LogError(message string, properties *map[string]string) {
	l.Log(dvb_logger.ERROR, message, properties)
}

func (l *Logger) LogFatal(message string, properties *map[string]string) {
	if properties == nil {
		l.logger.Fatalf(logFmtMessageFormat, dvb_logger.LevelNames[dvb_logger.FATAL], message)
	} else {
		l.logger.Fatalf(logFmtMessageFormatWithProps, mapPropertiesToLogFmt(*properties), dvb_logger.LevelNames[dvb_logger.FATAL], message)
	}
}

func (l *Logger) LogPanic(message string, properties *map[string]string) {
	if properties == nil {
		l.logger.Panicf(logFmtMessageFormat, dvb_logger.LevelNames[dvb_logger.FATAL], message)
	} else {
		l.logger.Panicf(logFmtMessageFormatWithProps, mapPropertiesToLogFmt(*properties), dvb_logger.LevelNames[dvb_logger.FATAL], message)
	}
}

func New(baseLogger *log.Logger, level int64, properties map[string]string) dvb_logger.Logger {
	baseLogger.SetPrefix(mapPropertiesToLogFmt(properties))
	return &Logger{
		level:  level,
		logger: baseLogger,
	}
}
