package rate_limiter

import (
	"gateway/internal/shared"
	"sync"
	"time"
)

type SlidingWindowRateLimiter struct {
	storage map[string]*slidingWindowRequestData
	config  shared.RateLimiterConfig
	mutex   sync.Mutex
}

type slidingWindowRequestData struct {
	timestamps []time.Time
}

func (rl *SlidingWindowRateLimiter) IsAllowed(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()

	data, exists := rl.storage[key]

	if !exists {
		rl.storage[key] = &slidingWindowRequestData{}
	} else {
		rl.cleanMap(data, now)
	}

	if len(rl.storage[key].timestamps) >= rl.config.Limit {
		return false
	}

	data.timestamps = append(data.timestamps, now)
	return true
}

func (rl *SlidingWindowRateLimiter) cleanMap(storage *slidingWindowRequestData, now time.Time) {
	window := time.Duration(rl.config.Rate) * time.Millisecond

	newSlice := storage.timestamps[:0]
	for _, timestamp := range storage.timestamps {
		if timestamp.After(now.Add(-window)) {
			newSlice = append(newSlice, timestamp)
		}
	}

	storage.timestamps = newSlice
}

func NewSlidingWindowRateLimiter(config shared.RateLimiterConfig) RateLimiter {
	return &SlidingWindowRateLimiter{
		storage: make(map[string]*slidingWindowRequestData),
		config:  config,
		mutex:   sync.Mutex{},
	}
}
