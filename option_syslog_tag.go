package log

type SyslogTagSetter interface {
	SetSyslogTag(tag string)
}

type SyslogTag struct {
	tag string
}

func (s *SyslogTag) SetSyslogTag(tag string) {
	s.tag = tag
}

type withSyslogTag string

func (w withSyslogTag) Apply(l Logger) error {
	if setter, ok := l.(SyslogTagSetter); ok {
		setter.SetSyslogTag(string(w))
	}

	return nil
}

// WithSyslogTag specifies the tag for syslog logger.
func WithSyslogTag(tag string) Option {
	return withSyslogTag(tag)
}

var _ Option = withSyslogTag("")
