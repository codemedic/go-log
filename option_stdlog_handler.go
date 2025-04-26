package log

import (
	"fmt"
	"io"
	stdlog "log"
)

type withStdlogHandler bool

func (w withStdlogHandler) Apply(l Logger) error {
	if lw, ok := l.(io.Writer); ok {
		stdlog.SetFlags(0)
		stdlog.SetOutput(lw)
		return nil
	}

	return fmt.Errorf("logger %T does not implement io.Writer; %w", l, ErrIncompatibleOption)
}

// WithStdlogHandler specifies whether the logger is to be setup as handler for logging through the standard logger. All
// messages that arrive via the global standard logger will be logged at INFO level.
func WithStdlogHandler(enable bool) Option {
	return withStdlogHandler(enable)
}

var _ Option = withStdlogHandler(false)
