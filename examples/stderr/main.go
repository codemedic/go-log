package main

import (
	"bytes"
	"log"
	"sync"

	golog "github.com/codemedic/go-log"
)

func main() {
	l := golog.Must(golog.NewStderr(
		golog.OptionsMust(golog.Options(
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			golog.WithUTCTimestampFromEnv("LOG_UTC", true),
			golog.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			golog.WithPrintLevel(golog.Info),
			golog.WithMicrosecondsTimestamp,
			golog.WithStdlogSorter(func(b []byte) golog.Level {
				if bytes.HasPrefix(b, []byte("DEBUG")) {
					return golog.Disabled
				}

				return golog.Info
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

	log.Print("DEBUG: hide me")
	log.Print("done")
}
