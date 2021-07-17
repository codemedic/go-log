package log

import "os"

type withLevel struct {
	level Level
}

func (w withLevel) applySyslog(l *syslogLogger) error {
	l.level = w.level
	return nil
}

func (w withLevel) applyStdLog(l *stdLevelLogger) error {
	l.level = w.level
	return nil
}

func WithLevel(level Level) Option {
	return withLevel{
		level: level,
	}
}

func WithLevelFromEnv(env string, defaultLevel Level) OptionLoader {
	return func() (Option, error) {
		level := defaultLevel
		if value, found := os.LookupEnv(env); found {
			var err error
			level, err = LevelFromString(value)
			if err != nil {
				return nil, invalidEnvValue(env, err)
			}
		}

		return WithLevel(level), nil
	}
}

var _ Option = withLevel{}
