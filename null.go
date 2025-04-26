package log

type nullLogger struct{}

func (n nullLogger) Level() Level                                    { return Disabled }
func (n nullLogger) PrintLevel() Level                               { return Disabled }
func (n nullLogger) Logf(_ Level, _ int, _ string, _ ...interface{}) {}
func (n nullLogger) Close()                                          {}
func (n nullLogger) Write(p []byte) (w int, err error)               { return len(p), nil }

var nulLogger nullLogger
var nullLog = Log{logger: &nulLogger}

// NewNullLog gives a null logger that does nothing. It is useful for testing or when you want to disable logging altogether.
// Using a zero initialized Log object would do most of the same, but this is a more explicit way to indicate that logging is disabled.
// Additionally, this also sets up standard logger to use the null logger.
func NewNullLog(opt ...Option) (log Log, err error) {
	// apply default options first
	if err = CommonOptions.Apply(nulLogger); err != nil {
		err = newConfigError(err)
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.Apply(nulLogger); err != nil {
			err = newConfigError(err)
			return
		}
	}

	return nullLog, nil
}
