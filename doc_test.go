package log_test

import (
	"errors"
	"math/rand"

	"github.com/codemedic/go-log"
)

func ExampleNewStderr() {
	l := log.Must(log.NewStderr(
		log.OptionMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithMicrosecondsTimestamp,
		))))
	defer l.Close()
}

func ExampleNewLogfile() {
	l := log.Must(log.NewLogfile("/tmp/test-logfile.log", 0644,
		log.OptionMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithMicrosecondsTimestamp,
		))))
	defer l.Close()
}

func ExampleNewSyslog() {
	l := log.Must(log.NewSyslog(
		log.OptionMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithSyslogTag("test-syslog"),
			// the default as provided by the standard library; this is just for demonstration
			log.WithSyslogDaemonURL("unixgram:///dev/log"),
		))))
	defer l.Close()
}

func Example() {
	l := log.Must(log.NewSyslog())
	defer l.Close()

	l.Debug("debug message")
	l.Debugf("formatted %s message", "debug")

	l.Info("informational message")
	l.Debugf("formatted %s message", "informational")

	l.Warning("warning message")
	l.Warningf("formatted %s message", "warning")

	l.Error("error message")
	l.Debugf("formatted %v message", errors.New("error"))

	if l.DebugEnabled() {
		// Deriving debug data can cost memory, cpu or both. Do it only if the data is
		// not going to be thrown away by the logger due to a higher threshold log-level.
		data := rand.Int()
		l.Debugf("data: %d", data)
	}
}
