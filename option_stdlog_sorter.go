package log

type logSorter func([]byte) Level

func defaultLogSorter(_ []byte) Level {
	return Info
}

type StdLogSorterSetter interface {
	SetStdLogSorter(logSorter)
}

type StdLogSorter struct {
	stdSorter logSorter
}

func (s *StdLogSorter) SetStdLogSorter(sorter logSorter) {
	s.stdSorter = sorter
}

type withStdlogSorter logSorter

func (w withStdlogSorter) Apply(l Logger) error {
	if setter, ok := l.(StdLogSorterSetter); ok {
		setter.SetStdLogSorter(logSorter(w))
	}

	return nil
}

// WithStdlogSorter sets up the callback that decides the level to which a log-line from global standard logger will be
// logged.
func WithStdlogSorter(sorter logSorter) Option {
	return withStdlogSorter(sorter)
}

var _ Option = withStdlogSorter(nil)
