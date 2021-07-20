package log

import (
	"io"
	"os"
)

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

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

// WithWriter specifies the writer for a logger.
func WithWriter(w io.WriteCloser) Option {
	// prevent the logger from closing stderr or stdout
	if w == os.Stderr || w == os.Stdout {
		w = nopCloser{w}
	}

	return withWriter{
		writer: w,
	}
}

var _ Option = withWriter{}
