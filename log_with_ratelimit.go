package log

import (
	"sync/atomic"
	"time"
)

type withRateLimitLogger struct {
	Logger
	rateLimit       int64
	periodSeconds   int64
	periodStartTime atomic.Int64
	count           atomic.Int64
	dropped         atomic.Int64
}

func (l *withRateLimitLogger) Logf(level Level, calldepth int, format string, value ...interface{}) {
	currentTime := time.Now().Unix()
	if currentTime-l.periodStartTime.Load() > l.periodSeconds {
		if l.periodStartTime.CompareAndSwap(l.periodStartTime.Load(), currentTime) {
			l.count.Store(0)
			l.dropped.Store(0)
		}
	}

	if l.count.Load() < l.rateLimit {
		l.count.Add(1)
		l.Logger.Logf(level, calldepth+1, format, value...)
	} else if l.dropped.Add(1) == 1 {
		l.Logger.Logf(Info, calldepth+1, "log rate limit exceeded: %d message(s) dropped in the last %d second(s)", l.rateLimit, l.periodSeconds)
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
