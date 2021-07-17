package log

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrIncompatibleOption = errors.New("incompatible option")
var ErrBadSyslogDaemonURL = errors.New("bad syslog daemon url")
var ErrUnknownOption = errors.New("unknown option")

type Option interface {
	StdLogOption
	SyslogOption
}

type OptionLoader func() (Option, error)

type options []Option

func (o options) applySyslog(l *SyslogLogger) error {
	for _, opt := range o {
		if err := opt.applySyslog(l); err != nil {
			return err
		}
	}

	return nil
}

func (o options) applyStdLog(l *StdLevelLogger) error {
	for _, opt := range o {
		if err := opt.applyStdLog(l); err != nil {
			return err
		}
	}

	return nil
}

var _ Option = options{}

func Options(opt ...interface{}) (Option, error) {
	opts := options{}
	for _, o := range opt {
		switch o := o.(type) {
		case Option:
			opts = append(opts, o)
		case OptionLoader:
			lo, err := o()
			if err != nil {
				return nil, err
			}

			opts = append(opts, lo)
		default:
			return nil, ErrUnknownOption
		}
	}

	return opts, nil
}

func OptionsMust(o Option, err error) Option {
	if err != nil {
		panic(fmt.Errorf("failed to load options; error:%w", err))
	}

	return o
}

func boolFromEnv(env string, defaultValue bool) (bool, error) {
	bo := defaultValue
	if value, found := os.LookupEnv(env); found {
		v, err := strconv.ParseBool(value)
		if err != nil {
			return bo, invalidEnvValue(env, err)
		}

		bo = v
	}

	return bo, nil
}

func invalidEnvValue(env string, err error) error {
	return fmt.Errorf("invalid value in env:%s; error:%w", env, err)
}
