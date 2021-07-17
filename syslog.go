package log

import (
	"fmt"
	stdlog "log"
	"log/syslog"
)

type SyslogOption interface {
	applySyslog(*SyslogLogger) error
}

type SyslogLogger struct {
	level   Level
	flags   flags
	tag     string
	addr    string
	network string
	loggers []*stdlog.Logger
	closers []func()
}

func (s *SyslogLogger) Close() {
	s.level = Disabled
	for i, closer := range s.closers {
		s.loggers[i] = nil
		closer()
	}
}

func (s *SyslogLogger) Level() Level {
	return s.level
}

func (s *SyslogLogger) Logf(level Level, format string, value ...interface{}) {
	if !level.IsEnabled(s.level) {
		return
	}

	if level > Error {
		level = Error
	}

	_ = s.loggers[level-s.level].Output(3, fmt.Sprintf(format, value...))
}

var syslogDefaultOptions, _ = Options(
	WithUTCTimestamp(false),
	WithMicrosecondsTimestamp(false),
	WithSourceLocation(true),
	WithLevel(Debug),
)

func NewSyslog(opt ...Option) (_ Log, err error) {
	l := &SyslogLogger{}

	// apply default options first
	if err = syslogDefaultOptions.applySyslog(l); err != nil {
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.applySyslog(l); err != nil {
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
			return
		}

		l.loggers = append(l.loggers, stdlog.New(w, "", int(l.flags)))
		l.closers = append(l.closers, func() { _ = w.Close() })
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
