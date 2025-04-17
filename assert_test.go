package log

import "testing"

func TestClearAssert(t *testing.T) {
	l := Must(NewAssertLog(
		WithLevel(Info),
		WithPrintLevel(Info),
	))
	defer l.Close()

	l.Debug("debug message")
	l.Debugf("formatted %s message", "debug")

	l.Info("informational message")
	l.Infof("formatted %s message", "informational")

	if !Assert(l).Contains(Debug, "debug message") {
		t.Errorf("expected debug message not found")
	}

	if !Assert(l).ContainsFormat(Debug, "formatted %s message") {
		t.Errorf("expected 'formatted %%s message' not found at %s level", Debug)
	}

	if !Assert(l).ContainsLevel(Info) {
		t.Errorf("expected info message not found")
	}

	ClearAssert(l)
	if Assert(l).ContainsLevel(Info) {
		t.Errorf("expected info message found after clear")
	}
}
