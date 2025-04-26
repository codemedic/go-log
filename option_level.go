package log

import "os"

type LevelSetter interface {
	SetLevel(level Level) error
}

type LevelledLogger struct {
	level Level
}

func (l *LevelledLogger) SetLevel(level Level) error {
	l.level = level
	return nil
}

type withLevel Level

func (w withLevel) Apply(l Logger) error {
	if setter, ok := l.(LevelSetter); ok {
		return setter.SetLevel(Level(w))
	}

	return nil
}

// WithLevel specifies the level for loggers.
func WithLevel(level Level) Option {
	return withLevel(level)
}

// WithLevelString specifies the level for loggers as a string.
func WithLevelString(str string) OptionLoader {
	return func() (Option, error) {
		level, err := levelFromString(str)
		if err != nil {
			return nil, newConfigError(err)
		}

		return withLevel(level), nil
	}
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

// ---------------------------------------------------------

type PrintLevelSetter interface {
	SetPrintLevel(level Level) error
}

type PrintLevelledLogger struct {
	printLevel Level
}

func (l *PrintLevelledLogger) SetPrintLevel(level Level) error {
	l.printLevel = level
	return nil
}

type withPrintLevel Level

func (w withPrintLevel) Apply(l Logger) error {
	if setter, ok := l.(PrintLevelSetter); ok {
		return setter.SetPrintLevel(Level(w))
	}

	return nil
}

// WithPrintLevel specifies the level for log.Print and log.Printf.
func WithPrintLevel(level Level) Option {
	return withPrintLevel(level)
}

var _ Option = withLevel(Debug)
var _ Option = withPrintLevel(Debug)
