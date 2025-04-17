package log

import (
	"sync"
)

type assertLogOption interface {
	applyAssertLog(logger *assertLogger) error
}

type assertLogger struct {
	level      Level
	printLevel Level
	msgs       AssertMsgs
	mu         sync.Mutex
}

func (a *assertLogger) Level() Level {
	return a.level
}

func (a *assertLogger) PrintLevel() Level {
	return a.printLevel
}

func (a *assertLogger) Logf(level Level, format string, value ...interface{}) {
	if a.level.IsEnabled(level) {
		a.mu.Lock()
		defer a.mu.Unlock()
		a.msgs = append(a.msgs, &assertMsg{
			level:  level,
			format: format,
			values: value,
		})
	}
}

func (a *assertLogger) Close() {
	a.level = Disabled
	a.printLevel = Disabled

	a.mu.Lock()
	defer a.mu.Unlock()
	a.msgs = nil
}

func NewAssertLog(opt ...Option) (log Log, err error) {
	l := &assertLogger{
		level:      Debug,
		printLevel: Debug,
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.applyAssertLog(l); err != nil {
			err = newConfigError(err)
			return
		}
	}

	return Log{logger: l}, nil
}
