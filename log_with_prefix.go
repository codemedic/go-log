package log

import "strings"

type PrefixLogger interface {
	PrefixLogf(level Level, calldepth int, prefix, format string, value ...interface{})
}

type withPrefixLogger struct {
	Logger
	prefix string
	logf   func(level Level, calldepth int, prefix, format string, value ...interface{})
}

func (l *withPrefixLogger) Logf(level Level, calldepth int, format string, value ...interface{}) {
	if l.logf == nil {
		return
	}

	l.logf(level, calldepth+1, l.prefix, format, value...)
}

// WithPrefix specifies a prefix for the logger.
func (l Log) WithPrefix(prefix string) Log {
	if l.logger == nil {
		return l
	}

	if prefix == "" {
		return l
	}

	// Escape the prefix to avoid it interfering with the format string; replace any % with %%
	prefix = strings.ReplaceAll(prefix, "%", "%%")

	// If the logger is already a withPrefixLogger, combine the prefixes. Also, use the logger within the withPrefixLogger
	parentLogger := l.logger
	if withPrefix, ok := l.logger.(*withPrefixLogger); ok {
		prefix = withPrefix.prefix + prefix
		parentLogger = withPrefix.Logger
	}

	// If the parentLogger is PrefixLogger
	if ppl, ok := parentLogger.(PrefixLogger); ok {
		l.logger = &withPrefixLogger{
			Logger: parentLogger,
			prefix: prefix,
			logf:   ppl.PrefixLogf,
		}
	} else {
		l.logger = &withPrefixLogger{
			Logger: parentLogger,
			prefix: prefix,
			logf: func(level Level, calldepth int, prefix, format string, value ...interface{}) {
				parentLogger.Logf(level, calldepth+1, prefix+format, value...)
			},
		}
	}

	return l
}
