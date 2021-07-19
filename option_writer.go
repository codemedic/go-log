package log

import "io"

type withWriter struct {
	writer io.WriteCloser
}

func (w withWriter) applySyslog(*syslogLogger) error {
	return ErrIncompatibleOption
}

func (w withWriter) applyStdLog(l *stdLogger) error {
	l.writer = w.writer
	return nil
}

// WithWriter specifies the writer for a logger.
//
// Example:
//   l, err := log.New(WithWriter(os.Stdout))
func WithWriter(w io.WriteCloser) Option {
	return withWriter{
		writer: w,
	}
}

var _ Option = withWriter{}
