package log

import (
	"os"
	"regexp"
)

type withSyslogDaemonURL string

func (w withSyslogDaemonURL) applySyslog(l *syslogLogger) error {
	if w == "" {
		l.network, l.addr = "", ""
		return nil
	}

	matches := syslogDaemonURLRegex.FindStringSubmatch(string(w))
	if len(matches) == 0 {
		return ErrBadSyslogDaemonURL
	}

	l.network, l.addr = matches[1], matches[2]
	return nil
}

func (w withSyslogDaemonURL) applyStdLog(*stdLogger) error {
	return ErrIncompatibleOption
}

// WithSyslogDaemonURL specifies the syslog daemon URL for syslog logger.
func WithSyslogDaemonURL(url string) Option {
	return withSyslogDaemonURL(url)
}

// WithSyslogDaemonURLFromEnv makes a WithSyslogDaemonURL option based on the specified environment variable env or
// defaultUrl if no environment variable was found.
func WithSyslogDaemonURLFromEnv(env, defaultUrl string) OptionLoader {
	return func() (Option, error) {
		url := defaultUrl
		if value, found := os.LookupEnv(env); found {
			if !syslogDaemonURLRegex.MatchString(value) {
				return nil, ErrBadSyslogDaemonURL
			}

			url = value
		}

		return withSyslogDaemonURL(url), nil
	}
}

var syslogDaemonURLRegex = regexp.MustCompile(`^(tcp[46]?|udp[46]?|unix(?:gram|packet)?)://([^:]+(?::\d)?)$`)

var _ Option = withSyslogDaemonURL("")
