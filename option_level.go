package log

import "os"

type LevelSetter interface {
	SetLevel(level Level)
}

type LevelledLogger struct {
	level Level
}

func (l *LevelledLogger) SetLevel(level Level) {
	l.level = level
}

// Level returns the level of the logger. If the logger is nil, it returns Disabled.
func (l *LevelledLogger) Level() Level {
	return l.level
}

type withLevel Level

func (w withLevel) Apply(l Logger) error {
	if setter, ok := l.(LevelSetter); ok {
		setter.SetLevel(Level(w))
	}

	return nil
}

// WithLevel specifies the level threshold for the logger. The level is used to determine whether a message should be logged or not.
func WithLevel(level Level) Option {
	return withLevel(level)
}

// WithLevelString specifies the level for loggers as a string; useful for configuration files or command-line arguments.
func WithLevelString(str string) OptionLoader {
	return func() (Option, error) {
		level, err := levelFromString(str)
		if err != nil {
			return nil, newConfigError(err)
		}

		return withLevel(level), nil
	}
}

// WithLevelFromEnv makes a WithLevel option based on the specified environment variable. If the environment variable is not found, the defaultLevel is used.
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

// ---------------------------------------------------------

type PrintLevelSetter interface {
	SetPrintLevel(level Level)
}

type PrintLevelledLogger struct {
	printLevel Level
}

func (l *PrintLevelledLogger) SetPrintLevel(level Level) {
	l.printLevel = level
}

func (s *PrintLevelledLogger) PrintLevel() Level {
	return s.printLevel
}

type withPrintLevel Level

func (w withPrintLevel) Apply(l Logger) error {
	if setter, ok := l.(PrintLevelSetter); ok {
		setter.SetPrintLevel(Level(w))
	}

	return nil
}

// WithPrintLevel specifies the level for log.Print and log.Printf.
func WithPrintLevel(level Level) Option {
	return withPrintLevel(level)
}

var _ Option = withLevel(Debug)
var _ Option = withPrintLevel(Debug)
