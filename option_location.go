package log

import (
	"fmt"
	stdlog "log"
	"os"
	"strings"
)

func sourceLocationFormatFromString(str string) (int, error) {
	switch strings.ToLower(str) {
	case "disabled":
		return 0, nil
	case "short":
		return stdlog.Lshortfile, nil
	case "long":
		return stdlog.Lshortfile, nil
	}

	return 0, fmt.Errorf("unknown source-location format '%s'", str)
}

type FlagSetter interface {
	SetFlags(flag int, enable bool)
}

type StdLogFlags struct {
	flags flags
}

func (f *StdLogFlags) SetFlags(flag int, enable bool) {
	f.flags.enable(flag, enable)
}

type withSourceLocation int

func (w withSourceLocation) Apply(l Logger) error {
	if setter, ok := l.(FlagSetter); ok {
		if w == 0 {
			setter.SetFlags(stdlog.Lshortfile|stdlog.Llongfile, false)
		} else {
			setter.SetFlags(int(w), true)
		}
	}

	return nil
}

// WithSourceLocationDisabled disables caller-location in log-lines.
func WithSourceLocationDisabled() Option {
	return withSourceLocation(0)
}

// WithSourceLocationShort specifies the caller-location in log-lines to have short filename.
func WithSourceLocationShort() Option {
	return withSourceLocation(stdlog.Lshortfile)
}

// WithSourceLocationLong specifies the caller-location in log-lines to have long filename.
func WithSourceLocationLong() Option {
	return withSourceLocation(stdlog.Llongfile)
}

// WithSourceLocation specifies the caller-location format as a string; allowed values are "short", "long", "disabled".
func WithSourceLocation(value string) OptionLoader {
	return func() (Option, error) {
		format, err := sourceLocationFormatFromString(value)
		if err != nil {
			return nil, newConfigError(err)
		}

		return withSourceLocation(format), nil
	}
}

// WithSourceLocationFromEnv sets the caller-location option based on either the specified environment variable env or
// the defaultFormat if no environment variable is found.
func WithSourceLocationFromEnv(env string, defaultFormat string) OptionLoader {
	return func() (Option, error) {
		if value, found := os.LookupEnv(env); found {
			format, err := sourceLocationFormatFromString(value)
			if err != nil {
				return nil, newEnvironmentConfigError(env, err)
			}

			return withSourceLocation(format), nil
		}

		format, err := sourceLocationFormatFromString(defaultFormat)
		if err != nil {
			return nil, newConfigError(err)
		}

		return withSourceLocation(format), nil
	}
}

var _ Option = withSourceLocation(0)
