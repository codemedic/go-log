package log

import "os"

type WithLevel Level

func (w WithLevel) applySyslog(l *syslogLogger) error {
	l.level = Level(w)
	return nil
}

func (w WithLevel) applyStdLog(l *stdLevelLogger) error {
	l.level = Level(w)
	return nil
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

var _ Option = WithLevel(Debug)
