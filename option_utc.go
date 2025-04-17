package log

import "log"

type withUTCTimestamp bool

func (w withUTCTimestamp) applyAssertLog(*assertLogger) error {
	return ErrIncompatibleOption
}

func (w withUTCTimestamp) applySyslog(l *syslogLogger) error {
	l.flags.enable(log.LUTC, bool(w))
	return nil
}

func (w withUTCTimestamp) applyStdLog(l *stdLogger) error {
	l.flags.enable(log.LUTC, bool(w))
	return nil
}

// WithUTCTimestamp specifies whether loggers are to log timestamp in UTC.
func WithUTCTimestamp(enable bool) Option {
	return withUTCTimestamp(enable)
}

// WithUTCTimestampFromEnv makes a WithUTCTimestamp option based on the specified environment variable env or
// defaultEnable if no environment variable was found.
func WithUTCTimestampFromEnv(env string, defaultEnable bool) OptionLoader {
	return func() (Option, error) {
		enable, err := boolFromEnv(env, defaultEnable)
		if err != nil {
			return nil, err
		}

		return withUTCTimestamp(enable), nil
	}
}

var _ Option = withUTCTimestamp(false)
