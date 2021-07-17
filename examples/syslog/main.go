package main

import (
	"github.com/codemedic/go-log"
	"sync"
)

func main() {
	l := log.Must(log.NewSyslog(
		log.OptionsMust(log.Options(
			log.WithLevelFromEnv("LOG_LEVEL", log.Info),
			log.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", true),
			log.WithSyslogTag("test-syslog"),
		))))
	defer l.Close()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			l.Debug("starting up...")
			for j := 0; j < 10000; j++ {
				l.Info("hello world")
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
