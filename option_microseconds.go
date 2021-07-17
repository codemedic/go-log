package log

import "log"

type WithMicrosecondsTimestamp bool

func (w WithMicrosecondsTimestamp) applySyslog(l *SyslogLogger) error {
	l.flags.enable(log.Lmicroseconds, bool(w))
	return nil
}

func (w WithMicrosecondsTimestamp) applyStdLog(l *StdLevelLogger) error {
	l.flags.enable(log.Lmicroseconds, bool(w))
	return nil
}

func WithMicrosecondsTimestampFromEnv(env string, defaultEnable bool) OptionLoader {
	return func() (Option, error) {
		enable, err := boolFromEnv(env, defaultEnable)
		if err != nil {
			return nil, err
		}

		return WithMicrosecondsTimestamp(enable), nil
	}
}

var _ Option = WithMicrosecondsTimestamp(false)
