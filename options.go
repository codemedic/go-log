package log

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	Level        Level
	UTC          bool
	Microseconds bool
	Location     bool
}

func newOptions() *Options {
	return &Options{
		Level:    Debug,
		UTC:      true,
		Location: false,
	}
}

var Opts struct {
	// Level sets the minimum level for the logger.
	Level func(level Level) func(opt *Options) error
	// LevelFromEnv sets Level based on value from a specified environment variable.
	// If no environment variable is found then defaultLevel is used.
	LevelFromEnv func(env string, defaultLevel Level) func(opt *Options) error

	// UTC enables or disables UTC timezone for the timestamp added to each line.
	UTC func(enable bool) func(opt *Options) error
	// UTCFromEnv enables or disables UTC based on value from a specified environment variable.
	// If no environment variable is found then defaultUTC is used.
	UTCFromEnv func(env string, defaultUTC bool) func(opt *Options) error

	// Microseconds enables or disables microseconds precision for the timestamp added to each line.
	Microseconds func(enable bool) func(opt *Options) error
	// MicrosecondsFromEnv enables or disables Microseconds based on value from a specified environment variable.
	// If no environment variable is found then defaultMicroseconds is used.
	MicrosecondsFromEnv func(env string, defaultMicroseconds bool) func(opt *Options) error

	// Location enables or disables source-location added to each line.
	Location func(enable bool) func(opt *Options) error
	// LocationFromEnv enables or disables Location based on value from a specified environment variable.
	// If no environment variable is found then defaultLocation is used.
	LocationFromEnv func(env string, defaultLocation bool) func(opt *Options) error
}

func getBoolEnv(name string, defaultValue bool) (bool, error) {
	if value, found := os.LookupEnv(name); found {
		return strconv.ParseBool(value)
	}

	return defaultValue, nil
}

func invalidEnvValue(env string, err error) error {
	return fmt.Errorf("invalid value in env:%s; error:%w", env, err)
}

func init() {
	Opts.Level = func(level Level) func(opt *Options) error {
		return func(opt *Options) error {
			opt.Level = level
			return nil
		}
	}

	Opts.LevelFromEnv = func(env string, defaultLevel Level) func(opt *Options) error {
		return func(opt *Options) error {
			level := defaultLevel
			if value, found := os.LookupEnv(env); found {
				var err error
				level, err = LevelFromString(value)
				if err != nil {
					return invalidEnvValue(env, err)
				}
			}

			opt.Level = level
			return nil
		}
	}

	Opts.UTC = func(enable bool) func(opt *Options) error {
		return func(opt *Options) error {
			opt.UTC = enable
			return nil
		}
	}

	Opts.UTCFromEnv = func(env string, defaultUTC bool) func(opt *Options) error {
		return func(opt *Options) error {
			enable, err := getBoolEnv(env, defaultUTC)
			if err != nil {
				return invalidEnvValue(env, err)
			}

			opt.UTC = enable
			return nil
		}
	}

	Opts.Microseconds = func(enable bool) func(opt *Options) error {
		return func(opt *Options) error {
			opt.Microseconds = enable
			return nil
		}
	}

	Opts.MicrosecondsFromEnv = func(env string, defaultUTC bool) func(opt *Options) error {
		return func(opt *Options) error {
			enable, err := getBoolEnv(env, defaultUTC)
			if err != nil {
				return invalidEnvValue(env, err)
			}

			opt.Microseconds = enable
			return nil
		}
	}

	Opts.Location = func(enable bool) func(opt *Options) error {
		return func(opt *Options) error {
			opt.Location = enable
			return nil
		}
	}

	Opts.LocationFromEnv = func(env string, defaultUTC bool) func(opt *Options) error {
		return func(opt *Options) error {
			enable, err := getBoolEnv(env, defaultUTC)
			if err != nil {
				return invalidEnvValue(env, err)
			}

			opt.Location = enable
			return nil
		}
	}
}
