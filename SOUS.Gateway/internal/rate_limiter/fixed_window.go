package rate_limiter

import (
	"gateway/internal/shared"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	storage map[string]*fixedWindowRequestData
	config  shared.RateLimiterConfig
	mutex   sync.Mutex
}

type fixedWindowRequestData struct {
	count    int
	earliest time.Time
}

func (rl *FixedWindowRateLimiter) IsAllowed(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	window := time.Duration(rl.config.Rate) * time.Millisecond

	data, exists := rl.storage[key]

	if exists && rl.storage[key].count >= rl.config.Limit {
		return false
	}

	if !exists || now.After(data.earliest.Add(window)) {
		rl.storage[key] = &fixedWindowRequestData{count: 0, earliest: now}
	}

	rl.storage[key].count++
	return true
}

// Cleanup TODO: Decouple storage from implementation, make it storage agnostic, with option of using smth like Redis.
// Cleanup won't be needed in Redis implementation, it has its own TTL mechanism
func (rl *FixedWindowRateLimiter) Cleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			rl.mutex.Lock()
			now := time.Now()
			window := time.Duration(rl.config.Rate) * time.Millisecond
			for key, data := range rl.storage {
				if now.After(data.earliest.Add(window)) {
					delete(rl.storage, key)
				}
			}

			rl.mutex.Unlock()
		}
	}()
}

func NewFixedWindowRateLimiter(config shared.RateLimiterConfig) RateLimiter {
	return &FixedWindowRateLimiter{
		storage: make(map[string]*fixedWindowRequestData),
		config:  config,
		mutex:   sync.Mutex{},
	}
}
