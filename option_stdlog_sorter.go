package log

type logSorter func([]byte) Level

func defaultLogSorter(_ []byte) Level {
	return Info
}

type withStdlogSorter logSorter

func (w withStdlogSorter) applyStdLog(l *stdLogger) error {
	if w != nil {
		l.stdSorter = logSorter(w)
	}

	return nil
}

func (w withStdlogSorter) applySyslog(l *syslogLogger) error {
	if w != nil {
		l.stdSorter = logSorter(w)
	}

	return nil
}

// WithStdlogSorter sets up the callback that decides the level to which a log-line from global standard logger will be
// logged.
func WithStdlogSorter(sorter logSorter) Option {
	return withStdlogSorter(sorter)
}

var _ Option = withStdlogSorter(nil)
