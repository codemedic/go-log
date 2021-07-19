package log

import (
	"fmt"
	stdlog "log"
	"os"
)

func sourceLocationFormatFromString(str string) (int, error) {
	switch str {
	case "DISABLED", "Disabled", "disabled":
		return 0, nil
	case "SHORT", "Short", "short":
		return stdlog.Lshortfile, nil
	case "LONG", "Long", "long":
		return stdlog.Lshortfile, nil
	}

	return 0, fmt.Errorf("unknown source-location format '%s'", str)
}

type withSourceLocation int

func (w withSourceLocation) applySyslog(l *syslogLogger) error {
	if w == 0 {
		l.flags.enable(stdlog.Lshortfile|stdlog.Llongfile, false)
	} else {
		l.flags.enable(int(w), true)
	}

	return nil
}

func (w withSourceLocation) applyStdLog(l *stdLogger) error {
	if w == 0 {
		l.flags.enable(stdlog.Lshortfile|stdlog.Llongfile, false)
	} else {
		l.flags.enable(int(w), true)
	}

	return nil
}

// WithSourceLocationDisabled disables caller-location in log-lines.
//
// Example:
//   l, err := log.NewSyslog(log.WithSourceLocationDisabled())
func WithSourceLocationDisabled() Option {
	return withSourceLocation(0)
}

// WithSourceLocationShort specifies the caller-location in log-lines to have short filename.
//
// Example:
//   l, err := log.NewSyslog(log.WithSourceLocationShort())
func WithSourceLocationShort() Option {
	return withSourceLocation(stdlog.Lshortfile)
}

// WithSourceLocationLong specifies the caller-location in log-lines to have long filename.
//
// Example:
//   l, err := log.NewSyslog(log.WithSourceLocationLong())
func WithSourceLocationLong() Option {
	return withSourceLocation(stdlog.Llongfile)
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
