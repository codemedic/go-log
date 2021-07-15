package log

type LevelLogger interface {
	Level() Level
	Logf(level Level, format string, value ...interface{})
}

type Log struct {
	logger LevelLogger
}

func (l Log) Debugf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Debug, format, value...)
}

func (l Log) Infof(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Info, format, value...)
}

func (l Log) Warningf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Warning, format, value...)
}

func (l Log) Errorf(format string, value ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Error, format, value...)
}

func (l Log) Debug(message string) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Debug, "%s", message)
}

func (l Log) Info(message string) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Info, "%s", message)
}

func (l Log) Warning(message string) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Warning, "%s", message)
}

func (l Log) Error(message string) {
	if l.logger == nil {
		return
	}

	l.logger.Logf(Error, "%s", message)
}

// DebugEnabled checks if DEBUG level is enabled for the logger.
// It can be used to check before performing any extra processing to generate data
// purely for logging, thereby avoiding the extra processing when DEBUG level is
// disabled.
//
// Example:
//     if logger.DebugEnabled() {
//         debugData := makeDebugData()
//         logger.Debugf("debug data: %v", debugData)
//     }
func (l Log) DebugEnabled() bool {
	if l.logger == nil {
		return false
	}

	return Debug.IsEnabled(l.logger.Level())
}
