package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
)

var DefaultLocationFormat = stdlog.Lshortfile

var stdLogDefaultOptions, _ = Options(
	WithUTCTimestamp(true),
	WithMicrosecondsTimestamp(true),
	WithSourceLocation(true),
	WithLevel(Debug),
	WithWriter(os.Stderr),
)

type stdLogOption interface {
	applyStdLog(*stdLevelLogger) error
}

type stdLevelLogger struct {
	level     Level
	flags     flags
	writer    io.WriteCloser
	stdLogger *stdlog.Logger
}

func (l *stdLevelLogger) Close() {
	l.level = Disabled
	_ = l.writer.Close()
}

func (l *stdLevelLogger) Level() Level {
	return l.level
}

func (l *stdLevelLogger) Logf(level Level, format string, value ...interface{}) {
	if level.IsEnabled(l.level) {
		_ = l.stdLogger.Output(3, fmt.Sprintf(level.String()+": "+format, value...))
	}
}

func NewStdLog(opt ...Option) (_ Log, err error) {
	l := &stdLevelLogger{
		flags: stdlog.LstdFlags,
	}

	// apply default options first
	if err = stdLogDefaultOptions.applyStdLog(l); err != nil {
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.applyStdLog(l); err != nil {
			return
		}
	}

	l.stdLogger = stdlog.New(l.writer, "", int(l.flags))

	return Log{logger: l}, nil
}
