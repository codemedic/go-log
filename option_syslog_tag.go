package log

type withSyslogTag string

func (w withSyslogTag) applyStdLog(*stdLogger) error {
	return ErrIncompatibleOption
}

func (w withSyslogTag) applySyslog(l *syslogLogger) error {
	l.tag = string(w)
	return nil
}

// WithSyslogTag specifies the tag for syslog logger.
func WithSyslogTag(tag string) Option {
	return withSyslogTag(tag)
}

var _ Option = withSyslogTag("")
