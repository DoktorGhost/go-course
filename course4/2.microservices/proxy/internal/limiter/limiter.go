package limiter

import (
	"sync"
	"time"
)

var ipLimits sync.Map

type IPLimiter struct {
	Limiter *limiter
}

// свой лимитер
type limiter struct {
	Count  int
	Start  time.Time
	Limit  int
	Period time.Duration
}

func newLimiter(limit int, period time.Duration) *limiter {
	startTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	return &limiter{0, startTime, limit, period}
}

func (l *limiter) Take() bool {
	currentTime := time.Now()

	if l.Count == 0 {
		l.Count++
		l.Start = currentTime
		return true
	}

	if l.Count < l.Limit {
		l.Count++
		return true
	}

	if currentTime.Sub(l.Start) < l.Period {
		return false
	}

	l.Count = 1
	l.Start = currentTime
	return true
}

func GetLimiterForIP(ip string) *IPLimiter {
	value, exists := ipLimits.Load(ip)
	if exists {
		return value.(*IPLimiter)
	}

	limit := &IPLimiter{
		Limiter: newLimiter(5, time.Minute),
	}
	ipLimits.Store(ip, limit)
	return limit
}
