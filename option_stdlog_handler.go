package log

import (
	"io"
	stdlog "log"
)

type withStdlogHandler bool

func (w withStdlogHandler) apply(l io.Writer) {
	stdlog.SetFlags(0)
	stdlog.SetOutput(l)
}

func (w withStdlogHandler) applyAssertLog(*assertLogger) error {
	return ErrIncompatibleOption
}

func (w withStdlogHandler) applySyslog(l *syslogLogger) error {
	w.apply(l)
	return nil
}

func (w withStdlogHandler) applyStdLog(l *stdLogger) error {
	w.apply(l)
	return nil
}

// WithStdlogHandler specifies whether the logger is to be setup as handler for logging through the standard logger. All
// messages that arrive via the global standard logger will be logged at INFO level.
func WithStdlogHandler(enable bool) Option {
	return withStdlogHandler(enable)
}

var _ Option = withStdlogHandler(false)
