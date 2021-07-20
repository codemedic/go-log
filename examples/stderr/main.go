package main

import (
	"bytes"
	stdlog "log"
	"sync"

	"github.com/codemedic/go-log"
)

func main() {
	l := log.Must(log.NewStderr(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithUTCTimestampFromEnv("LOG_UTC", true),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			log.WithPrintLevel(log.Info),
			log.WithMicrosecondsTimestamp,
			log.WithStdlogSorter(func(b []byte) log.Level {
				if bytes.HasPrefix(b, []byte("DEBUG")) {
					return log.Disabled
				}

				return log.Info
			}),
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

	stdlog.Print("DEBUG: hide me")
	stdlog.Print("done")
}
