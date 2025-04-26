package log

import (
	"os"
	"regexp"
)

type SyslogURLSetter interface {
	SetSyslogURL(network, addr string)
}

type SyslogURL struct {
	// Network is the network type (tcp, udp, unixgram, unixpacket).
	network string
	// Addr is the address of the syslog daemon.
	addr string
}

func (s *SyslogURL) SetSyslogURL(network, addr string) {
	s.network = network
	s.addr = addr
}

type withSyslogDaemonURL string

func (w withSyslogDaemonURL) Apply(l Logger) error {
	if setter, ok := l.(SyslogURLSetter); ok {
		network := ""
		addr := ""

		if w != "" {
			matches := syslogDaemonURLRegex.FindStringSubmatch(string(w))
			if len(matches) > 0 {
				network = matches[1]
				addr = matches[2]
			} else {
				return ErrBadSyslogDaemonURL
			}
		}

		setter.SetSyslogURL(network, addr)
	}

	return nil
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
