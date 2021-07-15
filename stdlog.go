package log

import (
	"fmt"
	stdlog "log"
)

var LocationFormat = stdlog.Lshortfile

type StdLevelLogger Level

func (l StdLevelLogger) Level() Level {
	return Level(l)
}

func (l StdLevelLogger) Logf(level Level, format string, value ...interface{}) {
	if level.IsEnabled(Level(l)) {
		_ = stdlog.Output(3, fmt.Sprintf(level.String()+": "+format, value...))
	}
}

func newStdLog(opt *Options) (LevelLogger, error) {
	flags := stdlog.LstdFlags

	if opt.UTC {
		flags |= stdlog.LUTC
	}

	if opt.Microseconds {
		flags |= stdlog.Lmicroseconds
	}

	if opt.Location {
		flags |= LocationFormat
	}

	stdlog.SetFlags(flags)

	return StdLevelLogger(opt.Level), nil
}
