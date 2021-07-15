package main

import (
	"github.com/codemedic/go-log"
)

func main() {
	l, err := log.NewStdLog(
		log.Opts.LevelFromEnv("LOG_LEVEL", log.Info),
		log.Opts.UTCFromEnv("LOG_UTC", true),
		log.Opts.LocationFromEnv("LOG_SOURCE_LOCATION", true),
		log.Opts.Microseconds(true),
	)
	if err != nil {
		panic(err)
	}

	l.Debug("starting up...")
	l.Info("hello world")
}
