package log

import (
	"sync"
)

type assertLogger struct {
	LevelledLogger
	PrintLevelledLogger
	msgs AssertMsgs
	mu   sync.Mutex
}

func (a *assertLogger) Logf(level Level, _ int, format string, value ...interface{}) {
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
	l := &assertLogger{}
	l.SetLevel(Debug)
	l.SetPrintLevel(Debug)

	// apply any specified options
	for _, o := range opt {
		if err = o.Apply(l); err != nil {
			err = newConfigError(err)
			return
		}
	}

	return Log{logger: l}, nil
}
