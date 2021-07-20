package log_test

import (
	"bytes"
	"errors"
	"math/rand"
	"os"

	golog "github.com/codemedic/go-log"
)

func ExampleNewStderr() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewStderr(
		golog.OptionsMust(golog.Options(
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			golog.WithUTCTimestampFromEnv("LOG_UTC", true),
			golog.WithSourceLocationDisabled,
			golog.WithMicrosecondsTimestamp,
		))))

	defer l.Close()
}

func ExampleNewStdout() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewStdout(
		golog.OptionsMust(golog.Options(
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			golog.WithUTCTimestampFromEnv("LOG_UTC", true),
			golog.WithSourceLocationLong,
			golog.WithMicrosecondsTimestamp,
		))))

	defer l.Close()
}

func ExampleNewLogfile() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewLogfile("/tmp/test-logfile.log", 0644,
		golog.OptionsMust(golog.Options(
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			golog.WithUTCTimestampFromEnv("LOG_UTC", true),
			golog.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			golog.WithMicrosecondsTimestamp,
		))))

	defer l.Close()
}

func ExampleNewSyslog() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.OptionsMust(golog.Options(
			// set the log-level dynamically from the environment
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			// set the syslog tag
			golog.WithSyslogTag("test-syslog"),
			// write to syslog server over UDP
			golog.WithSyslogDaemonURL("udp://syslog.acme.com:514"),
		))))

	defer l.Close()
}

func ExampleWithLevel() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithLevel(golog.Info),
	))

	defer l.Close()

	l.Debug("hide me")
	l.Info("hello world!")
}

func ExampleWithPrintLevel() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithPrintLevel(golog.Info),
	))

	defer l.Close()

	l.Print("hello world!")
}

func ExampleWithSourceLocationDisabled() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSourceLocationDisabled(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationLong() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSourceLocationLong(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationShort() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSourceLocationShort(),
	))

	defer l.Close()
}

func ExampleWithSourceLocationFromEnv() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(golog.OptionsMust(golog.Options(
		golog.WithSourceLocationFromEnv("LOG_CALLER_LOCATION", "short"),
	))))

	defer l.Close()
}

func ExampleWithMicrosecondsTimestamp() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithMicrosecondsTimestamp(true),
	))

	defer l.Close()
}

func ExampleWithMicrosecondsTimestampFromEnv() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(golog.OptionsMust(golog.Options(
		golog.WithMicrosecondsTimestampFromEnv("LOG_MICROSECOND_TIMESTAMP", true),
	))))

	defer l.Close()
}

func ExampleOptions() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.OptionsMust(
			golog.Options(
				golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
				golog.WithMicrosecondsTimestamp))))

	defer l.Close()
}

func ExampleOptionsMust() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.OptionsMust(
			golog.Options(
				golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
				golog.WithMicrosecondsTimestamp))))

	defer l.Close()
}

func ExampleWithStdlogHandler() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithStdlogHandler(false),
	))

	defer l.Close()
}

func ExampleWithStdlogSorter() {
	// import golog "github.com/codemedic/go-log"

	l, _ := golog.NewSyslog(golog.WithStdlogSorter(func(b []byte) golog.Level {
		switch {
		case bytes.HasPrefix(b, []byte("WARNING")):
			fallthrough
		case bytes.HasPrefix(b, []byte("ERROR")):
			return golog.Warning // ERROR and WARNING lines as Warning
		case bytes.HasPrefix(b, []byte("INFO")):
			fallthrough
		case bytes.HasPrefix(b, []byte("DEBUG")):
			return golog.Disabled // disable DEBUG & INFO lines
		default:
			return golog.Info // everything else as Info
		}
	}))

	defer l.Close()
}

func ExampleWithSyslogTag() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSyslogTag("my-app-name"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURL_uDP() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSyslogDaemonURL("udp://syslog.acme.com:514"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURL_local() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithSyslogDaemonURL("unixgram:///dev/log"),
	))

	defer l.Close()
}

func ExampleWithSyslogDaemonURLFromEnv() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(golog.OptionsMust(golog.Options(
		golog.WithSyslogDaemonURLFromEnv("LOG_SERVER", "udp://syslog.acme.com:514"),
	))))

	defer l.Close()
}

func ExampleWithUTCTimestamp() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithUTCTimestamp(true),
	))

	defer l.Close()
}

func ExampleWithUTCTimestampFromEnv() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(golog.OptionsMust(golog.Options(
		golog.WithUTCTimestampFromEnv("LOG_UTC", true),
	))))

	defer l.Close()
}

func ExampleWithWriter() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog(
		golog.WithWriter(os.Stdout),
	))

	defer l.Close()
}

func Example() {
	// import golog "github.com/codemedic/go-log"

	l := golog.Must(golog.NewSyslog())
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
