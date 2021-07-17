package log

type WithSyslogTag string

func (w WithSyslogTag) applyStdLog(*stdLevelLogger) error {
	return ErrIncompatibleOption
}

func (w WithSyslogTag) applySyslog(l *syslogLogger) error {
	l.tag = string(w)
	return nil
}

var _ Option = WithSyslogTag("")
