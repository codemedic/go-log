package log_test

import (
	"errors"
	"math/rand"

	"github.com/codemedic/go-log"
)

func ExampleNewStderr() {
	l := log.Must(log.NewStderr(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithMicrosecondsTimestamp,
		))))
	defer l.Close()
}

func ExampleNewLogfile() {
	l := log.Must(log.NewLogfile("/tmp/test-logfile.log", 0644,
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithMicrosecondsTimestamp,
		))))
	defer l.Close()
}

func ExampleNewSyslog() {
	l := log.Must(log.NewSyslog(
		log.OptionsMust(log.Options(
			// set the log-level dynamically from the environment
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			// set the syslog tag
			log.WithSyslogTag("test-syslog"),
			// write to syslog server over UDP
			log.WithSyslogDaemonURL("udp://syslog.acme.com:514"),
		))))
	defer l.Close()
}

func Example() {
	l := log.Must(log.NewSyslog())
	defer l.Close()

	l.Debug("debug message")
	l.Debugf("formatted %s message", "debug")

	l.Info("informational message")
	l.Infof("formatted %s message", "informational")

	l.Warning("warning message")
	l.Warningf("formatted %s message", "warning")

	l.Error("error message")
	l.Errorf("formatted %v message", errors.New("error"))

	if l.DebugEnabled() {
		// In cases where deriving debug data can be, costing memory, cpu or both, do it
		// only if the data is not going to be thrown away by the logger.
		data := rand.Int()
		l.Debugf("data: %d", data)
	}
}
