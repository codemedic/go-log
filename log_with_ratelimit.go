package log

import "time"

type withRateLimitLogger struct {
	Logger
	rateLimit       int64
	periodSeconds   int64
	periodStartTime int64
	count           int64
}

func (l *withRateLimitLogger) Logf(level Level, calldepth int, format string, value ...interface{}) {
	currentTime := time.Now().Unix()
	if currentTime-l.periodStartTime > l.periodSeconds {
		l.periodStartTime = currentTime
		l.count = 0
	}

	if l.count < l.rateLimit {
		l.count++
		l.Logger.Logf(level, calldepth+1, format, value...)
	}
}

// WithRateLimit creates a new logger with the specified rate limit.
func (l Log) WithRateLimit(rateLimit int64, periodSeconds int64) Log {
	if l.logger == nil {
		return l
	}

	// If the rate limit is disabled, return the original logger
	if rateLimit <= 0 {
		return l
	}

	if periodSeconds <= 0 {
		periodSeconds = 1
	}

	l.logger = &withRateLimitLogger{
		Logger:        l.logger,
		rateLimit:     rateLimit,
		periodSeconds: periodSeconds,
	}

	return l
}
