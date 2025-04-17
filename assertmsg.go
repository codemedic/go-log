package log

import "fmt"

type AssertMsg struct {
	level  Level
	format string
	values []interface{}
}

func (a *AssertMsg) Level() Level {
	return a.level
}
func (a *AssertMsg) Format() string {
	return a.format
}
func (a *AssertMsg) Values() []interface{} {
	return a.values
}
func (a *AssertMsg) Message() string {
	return fmt.Sprintf(a.format, a.values...)
}

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
		if m.level == level && m.Message() == msg {
			return true
		}
	}

	return false
}

func (a AssertMsgs) ContainsLevel(level Level) bool {
	for _, m := range a {
		if m.level == level {
			return true
		}
	}

	return false
}

func (a AssertMsgs) ContainsFormat(level Level, format string) bool {
	for _, m := range a {
		if m.level == level && m.format == format {
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
