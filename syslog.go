package log

import (
	"fmt"
	"io"
	stdlog "log"
	"log/syslog"
)

type syslogOption interface {
	applySyslog(*syslogLogger) error
}

type syslogLogger struct {
	level      Level
	flags      flags
	printLevel Level
	tag        string
	addr       string
	network    string
	loggers    []*stdlog.Logger
	closers    []func()
	stdHandler bool
	stdSorter  logSorter
}

// Write satisfies io.Writer interface so that syslogLogger can be used as writer for the standard global logger.
func (s *syslogLogger) Write(p []byte) (n int, err error) {
	level := s.stdSorter(p)
	logger := s.getLoggerByLevel(level)
	if logger == nil {
		return
	}

	err = logger.Output(4, string(p))
	return
}

func (s *syslogLogger) PrintLevel() Level {
	return s.printLevel
}

func (s *syslogLogger) Close() {
	s.level = Disabled
	for _, closer := range s.closers {
		closer()
	}
}

func (s *syslogLogger) Level() Level {
	return s.level
}

func (s *syslogLogger) getLoggerByLevel(level Level) *stdlog.Logger {
	if !level.IsEnabled(s.level) {
		return nil
	}

	if level > Error {
		level = Error
	}

	return s.loggers[level-s.level]
}

func (s *syslogLogger) Logf(level Level, format string, value ...interface{}) {
	logger := s.getLoggerByLevel(level)
	if logger == nil {
		return
	}

	_ = logger.Output(3, fmt.Sprintf(format, value...))
}

var syslogDefaultOptions, _ = Options(
	commonOptions,
	WithUTCTimestamp(false),
	WithMicrosecondsTimestamp(false),
)

// NewSyslog creates a new syslog logger with the specified options.
func NewSyslog(opt ...Option) (log Log, err error) {
	l := &syslogLogger{}

	// apply default options first
	if err = syslogDefaultOptions.applySyslog(l); err != nil {
		err = newConfigError(err)
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.applySyslog(l); err != nil {
			err = newConfigError(err)
			return
		}
	}

	if l.level == Disabled {
		return
	}

	for i := l.level; i <= Error; i++ {
		var w *syslog.Writer
		w, err = syslog.Dial(l.network, l.addr, toSyslogPriority(i), l.tag)
		if err != nil {
			err = newConnectionError(err)
			return
		}

		logger := stdlog.New(w, "", int(l.flags))
		l.loggers = append(l.loggers, logger)
		l.closers = append(l.closers, func() {
			// stop using the writer before closing it
			logger.SetOutput(io.Discard)
			_ = w.Close()
		})
	}

	return Log{logger: l}, nil
}

func toSyslogPriority(l Level) syslog.Priority {
	switch l {
	case Debug:
		return syslog.LOG_DEBUG
	case Warning:
		return syslog.LOG_WARNING
	case Error:
		return syslog.LOG_ERR
	case Info:
		fallthrough
	default:
		return syslog.LOG_INFO
	}
}
