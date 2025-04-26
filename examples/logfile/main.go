package main

import (
	"log"
	"sync"

	golog "github.com/codemedic/go-log"
)

func main() {
	l := golog.Must(golog.NewLogfile("test-logfile.log", 0644,
		golog.OptionsMust(golog.Options(
			golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
			golog.WithUTCTimestampFromEnv("LOG_UTC", true),
			golog.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
			golog.WithMicrosecondsTimestamp,
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

	log.Print("done")
}
