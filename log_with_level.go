package log

type withLevelLogger struct {
	Logger
	level Level
}

// Level returns the level of the logger. If the logger is nil, it returns Disabled.
func (l *withLevelLogger) Level() Level {
	return l.level
}

// Logf logs a message with the specified level and format. If the logger is nil, it does nothing.
func (l *withLevelLogger) Logf(level Level, calldepth int, format string, value ...interface{}) {
	if level.IsEnabled(l.level) {
		l.Logger.Logf(level, calldepth+1, format, value...)
	}
}

// WithLevel creates a new logger with the specified level. Note that the returned logger would still be limited to the level set on the original logger.
func (l Log) WithLevel(level Level) Log {
	if l.logger == nil {
		return l
	}

	// If the level is disabled, return the original logger
	if level.IsEnabled(l.logger.Level()) {
		l.logger = &withLevelLogger{
			Logger: l.logger,
			level:  level,
		}
	}

	return l
}
