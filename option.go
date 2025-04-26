package log

import (
	"os"
	"strconv"
)

// Option provides the interface through which all loggers can be configured.
type Option interface {
	Apply(l Logger) error
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

var CommonOptions, _ = Options(
	WithUTCTimestamp,
	WithMicrosecondsTimestamp,
	WithSourceLocationShort,
	WithLevel(Debug),
	WithPrintLevel(Info), // same level as the default for stdlog
	WithStdlogHandler,
	WithStdlogSorter(func(_ []byte) Level {
		return Info
	}),
)
