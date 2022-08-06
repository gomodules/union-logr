package ulogr

import "github.com/go-logr/logr"

func NewLogger(loggers ...logr.Logger) logr.Logger {
	out := make([]logr.LogSink, 0, len(loggers))
	for _, l := range loggers {
		sink := l.GetSink()
		if withCallDepth, ok := sink.(logr.CallDepthLogSink); ok {
			withCallDepth.WithCallDepth(2)
		}
		out = append(out, sink)
	}
	return logr.New(unionSink(out))
}

type unionSink []logr.LogSink

var _ logr.LogSink = unionSink{}

func (l unionSink) Init(info logr.RuntimeInfo) {
	for _, sink := range l {
		sink.Init(info)
	}
}

func (l unionSink) Enabled(level int) bool {
	enabled := false
	for _, sink := range l {
		enabled = enabled || sink.Enabled(level)
	}
	return enabled
}

func (l unionSink) Error(err error, msg string, keysAndValues ...interface{}) {
	for _, sink := range l {
		sink.Error(err, msg, keysAndValues...)
	}
}

func (l unionSink) Info(level int, msg string, keysAndValues ...interface{}) {
	for _, sink := range l {
		sink.Info(level, msg, keysAndValues...)
	}
}

func (l unionSink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	out := make([]logr.LogSink, 0, len(l))
	for _, sink := range l {
		out = append(out, sink.WithValues(keysAndValues...))
	}
	return unionSink(out)
}

func (l unionSink) WithName(name string) logr.LogSink {
	out := make([]logr.LogSink, 0, len(l))
	for _, sink := range l {
		out = append(out, sink.WithName(name))
	}
	return unionSink(out)
}
