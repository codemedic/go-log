package log

type WithSyslogTag string

func (w WithSyslogTag) applyStdLog(*StdLevelLogger) error {
	return ErrIncompatibleOption
}

func (w WithSyslogTag) applySyslog(l *SyslogLogger) error {
	l.tag = string(w)
	return nil
}

var _ Option = WithSyslogTag("")
