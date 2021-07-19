package log

type options []Option

func (o options) appendCopy(opt ...Option) options {
	n := options{}
	n.append(o...)
	n.append(opt...)

	return n
}

func (o *options) append(opt ...Option) {
	*o = append(*o, opt...)
}

func (o options) applySyslog(l *syslogLogger) error {
	for _, opt := range o {
		if err := opt.applySyslog(l); err != nil {
			return err
		}
	}

	return nil
}

func (o options) applyStdLog(l *stdLogger) error {
	for _, opt := range o {
		if err := opt.applyStdLog(l); err != nil {
			return err
		}
	}

	return nil
}

type OptionLoader func() (Option, error)

// Options combine multiple options into one composite option. It takes a list of Option or OptionLoader; the latter
// makes it possible to load options dynamically from environment, config files, etc
func Options(opt ...interface{}) (Option, error) {
	opts := options{}
	for _, o := range opt {
		switch o := o.(type) {
		case Option:
			opts.append(o)
		case OptionLoader:
			lo, err := o()
			if err != nil {
				return nil, err
			}

			opts.append(lo)
		case func() Option:
			opts.append(o())
		case func(bool) Option:
			opts.append(o(true))
		default:
			return nil, ErrUnknownOption
		}
	}

	return opts, nil
}

var _ Option = options{}
