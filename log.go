package log

import (
	"fmt"
	"os"
)

const defaultCallDepth = 3

// Logger interface provides means to extend this library
type Logger interface {
	// Level gives the current threshold of the logger
	Level() Level
	// PrintLevel gives the level at which log.Print logs
	PrintLevel() Level
	// Logf is the workhorse function that logs each line; works in a similar way to fmt.Printf
	Logf(level Level, calldepth int, format string, value ...interface{})
	// Close closes the logger
	Close()
}

type Log struct {
	logger Logger
}

func (l Log) Level() Level {
	if l.logger == nil {
		return Disabled
	}

	return l.logger.Level()
}

func (l Log) Debugf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Debug, defaultCallDepth, format, value...)
}

func (l Log) Infof(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Info, defaultCallDepth, format, value...)
}

func (l Log) Warningf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Warning, defaultCallDepth, format, value...)
}

func (l Log) Errorf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Error, defaultCallDepth, format, value...)
}

func (l Log) Printf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(l.logger.PrintLevel(), defaultCallDepth, format, value...)
}

func (l Log) Fatalf(format string, value ...interface{}) {
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, format, value...)
	}

	os.Exit(1)
}

func (l Log) Debug(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Debug, defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Info(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Info, defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Warning(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Warning, defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Error(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Error, defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Print(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(l.logger.PrintLevel(), defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Println(value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(l.logger.PrintLevel(), defaultCallDepth, "%s", fmt.Sprint(value...))
}

func (l Log) Fatal(value ...interface{}) {
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, "%s", fmt.Sprint(value...))
	}

	os.Exit(1)
}

func (l Log) Fatalln(value ...interface{}) {
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, "%s", fmt.Sprint(value...))
	}

	os.Exit(1)
}

func (l Log) Panic(value ...interface{}) {
	s := fmt.Sprint(value...)
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, "%s", s)
	}

	panic(s)
}

func (l Log) Panicf(format string, value ...interface{}) {
	s := fmt.Sprintf(format, value...)
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, "%s", s)
	}

	panic(s)
}

func (l Log) Panicln(value ...interface{}) {
	s := fmt.Sprint(value...)
	if l.logger != nil {
		l.logger.Logf(Error, defaultCallDepth, "%s", s)
	}

	panic(s)
}

// DebugEnabled checks if DEBUG level is enabled for the logger.
// It can be used to check before performing any extra processing to generate data
// that is purely for logging, thereby avoiding the extra processing when DEBUG
// level is disabled.
//
// Example:
//
//	if logger.DebugEnabled() {
//	  debugData := makeDebugData()
//	  logger.Debugf("debug data: %v", debugData)
//	}
func (l Log) DebugEnabled() bool {
	if l.logger == nil {
		return false
	}

	return Debug.IsEnabled(l.logger.Level())
}

// Close disables and closes the logger, freeing up any resources allocated to the logger.
// Once closed the logger will be disabled, but it will remain safe to use (free from panics).
func (l Log) Close() {
	if l.logger != nil {
		l.logger.Close()
	}
}

// Must ensures that a Log instance was initialised without error; panics if there was an error.
func Must(l Log, err error) Log {
	if err != nil {
		panic(fmt.Errorf("failed to initialise logger; %w", err))
	}

	return l
}
