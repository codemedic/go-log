package log

import "os"

type withLevel Level

func (w withLevel) applySyslog(l *syslogLogger) error {
	l.level = Level(w)
	return nil
}

func (w withLevel) applyStdLog(l *stdLogger) error {
	l.level = Level(w)
	return nil
}

// WithLevel specifies the level for loggers.
//
// Example:
//   l, err := log.NewSyslog(WithLevel(Info))
func WithLevel(level Level) Option {
	return withLevel(level)
}

// WithLevelFromEnv makes a WithLevel option based on the specified environment variable env or defaultLevel if no
// environment variable was found.
func WithLevelFromEnv(env string, defaultLevel Level) OptionLoader {
	return func() (Option, error) {
		level := defaultLevel
		if value, found := os.LookupEnv(env); found {
			var err error
			level, err = levelFromString(value)
			if err != nil {
				return nil, newEnvironmentConfigError(env, err)
			}
		}

		return withLevel(level), nil
	}
}

var _ Option = withLevel(Debug)
