package log

import "log"

type withMicrosecondsTimestamp bool

func (w withMicrosecondsTimestamp) Apply(l Logger) error {
	if setter, ok := l.(FlagSetter); ok {
		setter.SetFlags(log.Lmicroseconds, bool(w))
	}

	return nil
}

// WithMicrosecondsTimestamp specifies whether loggers are to log timestamp with microseconds precision.
func WithMicrosecondsTimestamp(enable bool) Option {
	return withMicrosecondsTimestamp(enable)
}

// WithMicrosecondsTimestampFromEnv makes a WithMicrosecondsTimestamp option based on the specified environment variable
// env or defaultEnable if no environment variable was found.
func WithMicrosecondsTimestampFromEnv(env string, defaultEnable bool) OptionLoader {
	return func() (Option, error) {
		enable, err := boolFromEnv(env, defaultEnable)
		if err != nil {
			return nil, err
		}

		return withMicrosecondsTimestamp(enable), nil
	}
}

var _ Option = withMicrosecondsTimestamp(false)
