package ulogr

import (
	"github.com/go-logr/logr"
)

type unionLogger struct {
	loggers     []logr.Logger
	infoLoggers []logr.InfoLogger
}

var _ logr.Logger = unionLogger{}

func NewUnionLogger(loggers ...logr.Logger) logr.Logger {
	return unionLogger{loggers: loggers}
}

func (l unionLogger) Info(msg string, keysAndValues ...interface{}) {
	for _, logger := range l.loggers {
		logger.Info(msg, keysAndValues...)
	}

	for _, infoLogger := range l.infoLoggers {
		infoLogger.Info(msg, keysAndValues...)
	}
}

func (l unionLogger) Enabled() bool {
	return true
}

func (l unionLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	for _, logger := range l.loggers {
		logger.Error(err, msg, keysAndValues...)
	}
}

func (l unionLogger) V(level int) logr.InfoLogger {
	out := make([]logr.InfoLogger, 0)
	for _, logger := range l.loggers {
		out = append(out, logger.V(level))
	}

	return unionLogger{infoLoggers: out}
}

func (l unionLogger) WithName(name string) logr.Logger {
	out := make([]logr.Logger, 0)
	for _, logger := range l.loggers {
		out = append(out, logger.WithName(name))
	}
	return unionLogger{loggers: out}
}

func (l unionLogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	out := make([]logr.Logger, len(l.loggers))
	for _, logger := range l.loggers {
		out = append(out, logger.WithValues(keysAndValues...))
	}
	return unionLogger{loggers: out}
}
