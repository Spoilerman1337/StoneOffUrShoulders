package rate_limiter

import (
	"gateway/internal/shared"
	"sync"
	"time"
)

type TokenBucketRateLimiter struct {
	token       int
	config      shared.RateLimiterConfig
	mutex       sync.Mutex
	lastUpdated time.Time
}

func (rl *TokenBucketRateLimiter) IsAllowed(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastUpdated)

	if rl.token == 0 {
		return false
	}

	rl.token = rl.token - 1
	newTokens := int(float64(elapsed) / float64(rl.config.Rate) * float64(rl.config.TokensPerRate))

	if rl.token+newTokens >= rl.config.Limit {
		rl.token = rl.config.Limit
	} else {
		rl.token = rl.token + newTokens
	}

	rl.lastUpdated = now

	return true
}

func NewTokenBucketRateLimiter(config shared.RateLimiterConfig) RateLimiter {
	return &TokenBucketRateLimiter{
		token:  0,
		config: config,
		mutex:  sync.Mutex{},
	}
}
