package log

import "fmt"

type options []Option

func (o options) Apply(l Logger) error {
	for _, opt := range o {
		if err := opt.Apply(l); err != nil {
			return err
		}
	}

	return nil
}

func (o *options) append(opt ...Option) {
	*o = append(*o, opt...)
}

type OptionLoader func() (Option, error)

// Options combine multiple options into one composite option. It takes a list of Option or OptionLoader; the latter
// makes it possible to load options dynamically from environment, config files, etc
func Options(opt ...interface{}) (Option, error) {
	opts := options{}
	for _, o := range opt {
		switch o := o.(type) {
		case Option: // static option
			opts.append(o)
		case OptionLoader: // dynamic option loader
			lo, err := o()
			if err != nil {
				return nil, err
			}
			opts.append(lo)
		case func() Option: // static option loader
			opts.append(o())
		case func(bool) Option: // static option loader with boolean
			opts.append(o(true))
		default:
			return nil, ErrUnknownOption
		}
	}

	return opts, nil
}

// OptionsMust checks for errors from dynamic OptionLoader combined through Options.
// It panics if err is not nil otherwise returns o.
func OptionsMust(o Option, err error) Option {
	if err != nil {
		panic(fmt.Errorf("failed to load options; error:%w", err))
	}

	return o
}

var _ Option = options{}
