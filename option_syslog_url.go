package log

import (
	"os"
	"regexp"
)

type WithSyslogDaemonURL string

func (w WithSyslogDaemonURL) applySyslog(l *syslogLogger) error {
	if w == "" {
		l.network, l.addr = "", ""
		return nil
	}

	matches := syslogDaemonURLRegex.FindStringSubmatch(string(w))
	if len(matches) == 0 {
		return ErrBadSyslogDaemonURL
	}

	l.network, l.addr = matches[0], matches[1]
	return nil
}

func (w WithSyslogDaemonURL) applyStdLog(*stdLevelLogger) error {
	return ErrIncompatibleOption
}

func WithSyslogDaemonURLFromEnv(env, defaultUrl string) OptionLoader {
	return func() (Option, error) {
		url := defaultUrl
		if value, found := os.LookupEnv(env); found {
			if !syslogDaemonURLRegex.MatchString(value) {
				return nil, ErrBadSyslogDaemonURL
			}

			url = value
		}

		return WithSyslogDaemonURL(url), nil
	}
}

var syslogDaemonURLRegex = regexp.MustCompile(`^(tcp[46]?|udp[46]?|unix(?:gram|packet)?)://([^:]+:\d)$`)

var _ Option = WithSyslogDaemonURL("")
