package main

import (
	stdlog "log"
	"sync"

	"github.com/codemedic/go-log"
)

func main() {
	l := log.Must(log.NewSyslog(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithSyslogTag("test-syslog"),
			// the default as provided by the standard library; this is just for demonstration
			log.WithSyslogDaemonURL("unixgram:///dev/log"),
		))))
	defer l.Close()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			l.Debug("starting up...")
			for j := 0; j < 100; j++ {
				l.Info("hello world")
			}
			wg.Done()
		}()
	}

	wg.Wait()

	stdlog.Print("done")
}
