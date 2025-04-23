package main

import (
	golog "github.com/codemedic/go-log"
	"sync"
)

func runModule(l golog.Log) {
	l.Info("started")
	l.Debug("debug message")
	l.Warning("warning message")
	l.Error("error message")
}

func main() {
	l := golog.Must(golog.NewStderr(
		golog.OptionsMust(golog.Options(
			golog.WithLevel(golog.Debug),
			golog.WithUTCTimestamp,
			golog.WithMicrosecondsTimestamp,
			golog.WithSourceLocationShort(),
			golog.WithPrintLevel(golog.Debug),
		))))
	defer l.Close()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		runModule(l.WithPrefix("module 1 "))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		runModule(l.WithLevel(golog.Warning).WithPrefix("module 2 "))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		runModule(l.WithPrefix("module 3 ").WithLevel(golog.Error))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		runModule(l.WithPrefix("module 4 ").WithLevel(golog.Error).WithLevel(golog.Info))
		wg.Done()
	}()

	wg.Wait()
}
