package log

import "fmt"

type AssertMsg interface {
	Level() Level
	Format() string
	Values() []interface{}
	Message() string
}

type assertMsg struct {
	level  Level
	format string
	values []interface{}
}

func (a *assertMsg) Level() Level {
	return a.level
}
func (a *assertMsg) Format() string {
	return a.format
}
func (a *assertMsg) Values() []interface{} {
	return a.values
}
func (a *assertMsg) Message() string {
	return fmt.Sprintf(a.format, a.values...)
}
func (a *assertMsg) locked() {}

type AssertMsgs []AssertMsg

func Assert(l Log) AssertMsgs {
	if l.logger == nil {
		return nil
	}

	al, ok := l.logger.(*assertLogger)
	if !ok {
		return nil
	}

	al.mu.Lock()
	defer al.mu.Unlock()

	return al.msgs
}

func (a AssertMsgs) Contains(level Level, msg string) bool {
	for _, m := range a {
		if m.Level() == level && m.Message() == msg {
			return true
		}
	}

	return false
}

func (a AssertMsgs) ContainsLevel(level Level) bool {
	for _, m := range a {
		if m.Level() == level {
			return true
		}
	}

	return false
}

func (a AssertMsgs) ContainsFormat(level Level, format string) bool {
	for _, m := range a {
		if m.Level() == level && m.Format() == format {
			return true
		}
	}

	return false
}

func ClearAssert(l Log) {
	if l.logger == nil {
		return
	}

	al, ok := l.logger.(*assertLogger)
	if !ok {
		return
	}

	al.mu.Lock()
	defer al.mu.Unlock()

	al.msgs = nil
}
