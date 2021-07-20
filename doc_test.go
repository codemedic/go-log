package log_test

import (
	"bytes"
	"errors"
	"math/rand"
	"os"

	"github.com/codemedic/go-log"
)

func ExampleNewStderr() {
	l := log.Must(log.NewStderr(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationDisabled,
			log.WithMicrosecondsTimestamp,
		))))

	defer l.Close()
}

func ExampleNewStdout() {
	l := log.Must(log.NewStdout(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationLong,
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

func ExampleWithLevel() {
	l := log.Must(log.NewSyslog(
		log.WithLevel(log.Info),
	))

	defer l.Close()

	l.Debug("hide me")
	l.Info("hello world!")
}

func ExampleWithPrintLevel() {
	l := log.Must(log.NewSyslog(
		log.WithPrintLevel(log.Info),
	))

	defer l.Close()

	l.Print("hello world!")
}

func ExampleWithSourceLocationDisabled() {
	l := log.Must(log.NewSyslog(
		log.WithSourceLocationDisabled(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationLong() {
	l := log.Must(log.NewSyslog(
		log.WithSourceLocationLong(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationShort() {
	l := log.Must(log.NewSyslog(
		log.WithSourceLocationShort(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationFromEnv() {
	l := log.Must(log.NewSyslog(log.OptionsMust(log.Options(
		log.WithSourceLocationFromEnv("LOG_CALLER_LOCATION", "short"),
	))))

	defer l.Close()
}

func ExampleWithMicrosecondsTimestamp() {
	l := log.Must(log.NewSyslog(
		log.WithMicrosecondsTimestamp(true),
	))

	defer l.Close()
}

func ExampleWithMicrosecondsTimestampFromEnv() {
	l := log.Must(log.NewSyslog(log.OptionsMust(log.Options(
		log.WithMicrosecondsTimestampFromEnv("LOG_MICROSECOND_TIMESTAMP", true),
	))))

	defer l.Close()
}

func ExampleOptions() {
	l := log.Must(log.NewSyslog(
		log.OptionsMust(
			log.Options(
				log.WithLevelFromEnv("LOG_LEVEL", log.Info),
				log.WithMicrosecondsTimestamp))))

	defer l.Close()
}

func ExampleOptionsMust() {
	l := log.Must(log.NewSyslog(
		log.OptionsMust(
			log.Options(
				log.WithLevelFromEnv("LOG_LEVEL", log.Info),
				log.WithMicrosecondsTimestamp))))

	defer l.Close()
}

func ExampleWithStdlogHandler() {
	l := log.Must(log.NewSyslog(
		log.WithStdlogHandler(false),
	))

	defer l.Close()
}

func ExampleWithStdlogSorter() {
	l, _ := log.NewSyslog(log.WithStdlogSorter(func(b []byte) log.Level {
		switch {
		case bytes.HasPrefix(b, []byte("WARNING")):
			fallthrough
		case bytes.HasPrefix(b, []byte("ERROR")):
			return log.Warning // ERROR and WARNING lines as Warning
		case bytes.HasPrefix(b, []byte("INFO")):
			fallthrough
		case bytes.HasPrefix(b, []byte("DEBUG")):
			return log.Disabled // disable DEBUG & INFO lines
		default:
			return log.Info // everything else as Info
		}
	}))

	defer l.Close()
}

func ExampleWithSyslogTag() {
	l := log.Must(log.NewSyslog(
		log.WithSyslogTag("my-app-name"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURL_udp() {
	l := log.Must(log.NewSyslog(
		log.WithSyslogDaemonURL("udp://syslog.acme.com:514"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURL_local() {
	l := log.Must(log.NewSyslog(
		log.WithSyslogDaemonURL("unixgram:///dev/log"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURLFromEnv() {
	l := log.Must(log.NewSyslog(log.OptionsMust(log.Options(
		log.WithSyslogDaemonURLFromEnv("LOG_SERVER", "udp://syslog.acme.com:514"),
	))))

	defer l.Close()
}

func ExampleWithUTCTimestamp() {
	l := log.Must(log.NewSyslog(
		log.WithUTCTimestamp(true),
	))

	defer l.Close()
}

func ExampleWithUTCTimestampFromEnv() {
	l := log.Must(log.NewSyslog(log.OptionsMust(log.Options(
		log.WithUTCTimestampFromEnv("LOG_UTC", true),
	))))

	defer l.Close()
}

func ExampleWithWriter() {
	l := log.Must(log.NewSyslog(
		log.WithWriter(os.Stdout),
	))

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

	// In cases where deriving debug data has a significant cost to memory, cpu or both, do it
	// only if the data is not going to be thrown away by the logger.
	if l.DebugEnabled() {
		data := rand.Int()
		l.Debugf("data: %d", data)
	}
}
