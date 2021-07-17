package log

type WithSourceLocation bool

func (w WithSourceLocation) applySyslog(l *syslogLogger) error {
	l.flags.enable(DefaultLocationFormat, bool(w))
	return nil
}

func (w WithSourceLocation) applyStdLog(l *stdLevelLogger) error {
	l.flags.enable(DefaultLocationFormat, bool(w))
	return nil
}

func WithSourceLocationFromEnv(env string, defaultEnable bool) OptionLoader {
	return func() (Option, error) {
		enable, err := boolFromEnv(env, defaultEnable)
		if err != nil {
			return nil, err
		}

		return WithSourceLocation(enable), nil
	}
}

var _ Option = WithSourceLocation(false)
