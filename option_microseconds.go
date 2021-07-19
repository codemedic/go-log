package log

import "log"

type withMicrosecondsTimestamp bool

func (w withMicrosecondsTimestamp) applySyslog(l *syslogLogger) error {
	l.flags.enable(log.Lmicroseconds, bool(w))
	return nil
}

func (w withMicrosecondsTimestamp) applyStdLog(l *stdLogger) error {
	l.flags.enable(log.Lmicroseconds, bool(w))
	return nil
}

// WithMicrosecondsTimestamp specifies whether loggers are to log timestamp with microseconds precision.
//
// Example:
//   l, err := log.NewStderr(WithMicrosecondsTimestamp(true))
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
