package log

import (
	"fmt"
	"os"
	"strconv"
)

// Option provides the interface through which all loggers can be configured.
type Option interface {
	stdLogOption
	syslogOption
}

// OptionMust checks for err to be not nil and panics. If err is nil o is returned.
func OptionMust(o Option, err error) Option {
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
			return bo, newEnvironmentConfigError(env, err)
		}

		bo = v
	}

	return bo, nil
}
