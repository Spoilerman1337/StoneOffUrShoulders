package rate_limiter

import (
	"gateway/internal/shared"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	storage map[string]*requestData
	config  shared.RateLimiterConfig
	mutex   sync.Mutex
}

type requestData struct {
	count    int
	earliest time.Time
}

func (rl *FixedWindowRateLimiter) IsAllowed(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	window := time.Duration(rl.config.Rate)

	data, exists := rl.storage[key]

	if exists && rl.storage[key].count >= rl.config.Limit {
		return false
	}

	if !exists || now.After(data.earliest.Add(window)) {
		rl.storage[key] = &requestData{count: 0, earliest: now}
	}

	rl.storage[key].count++
	return true
}
