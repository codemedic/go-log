package log

import "log"

type WithUTCTimestamp bool

func (w WithUTCTimestamp) applySyslog(l *SyslogLogger) error {
	l.flags.enable(log.LUTC, bool(w))
	return nil
}

func (w WithUTCTimestamp) applyStdLog(l *StdLevelLogger) error {
	l.flags.enable(log.LUTC, bool(w))
	return nil
}

func WithUTCTimestampFromEnv(env string, defaultEnable bool) OptionLoader {
	return func() (Option, error) {
		enable, err := boolFromEnv(env, defaultEnable)
		if err != nil {
			return nil, err
		}

		return WithUTCTimestamp(enable), nil
	}
}

var _ Option = WithUTCTimestamp(false)
