package log

import "io"

type withWriter struct {
	writer io.WriteCloser
}

func (w withWriter) applySyslog(*SyslogLogger) error {
	return ErrIncompatibleOption
}

func (w withWriter) applyStdLog(l *StdLevelLogger) error {
	l.writer = w.writer
	return nil
}

func WithWriter(w io.WriteCloser) Option {
	return withWriter{
		writer: w,
	}
}

var _ Option = withWriter{}
