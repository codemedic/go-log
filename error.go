package log

import (
	"errors"
	"fmt"
)

// ErrIncompatibleOption occurs when any of the given options are not compatible with the object being configured.
var ErrIncompatibleOption = errors.New("incompatible option")

// ErrBadSyslogDaemonURL occurs when a given URL is not valid for syslog initialisation.
var ErrBadSyslogDaemonURL = errors.New("bad syslog daemon url")

// ErrUnknownOption signifies an option that couldn't be recognised.
var ErrUnknownOption = errors.New("unknown option")

// ErrBadLevel signifies that a string couldn't be translated to a valid logging level.
var ErrBadLevel = errors.New("bad level")

type wrappedError struct {
	err error
}

func (w *wrappedError) As(target interface{}) bool {
	//goland:noinspection GoErrorsAs
	return errors.As(w.err, target)
}

func (w *wrappedError) Is(target error) bool {
	return errors.Is(w.err, target)
}

func (w *wrappedError) Unwrap() error {
	return w.err
}

type ConfigError struct {
	wrappedError
}

func (c *ConfigError) Error() string {
	return fmt.Sprintf("configuration error; %s", c.err.Error())
}

func newConfigError(err error) error {
	return &ConfigError{
		wrappedError{
			err: err,
		},
	}
}

func newEnvironmentConfigError(env string, err error) error {
	return fmt.Errorf("bad value in environment variable %s; %w", env, newConfigError(err))
}

type ConnectionError struct {
	wrappedError
}

func (c *ConnectionError) Error() string {
	return fmt.Sprintf("connection error; error:%s", c.err.Error())
}

func newConnectionError(err error) error {
	return &ConnectionError{
		wrappedError{
			err: err,
		},
	}
}
