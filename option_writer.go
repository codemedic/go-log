package log

import "io"

type withWriter struct {
	writer io.WriteCloser
}

func (w withWriter) applySyslog(*syslogLogger) error {
	return ErrIncompatibleOption
}

func (w withWriter) applyStdLog(l *stdLevelLogger) error {
	l.writer = w.writer
	return nil
}

func WithWriter(w io.WriteCloser) Option {
	return withWriter{
		writer: w,
	}
}

var _ Option = withWriter{}
