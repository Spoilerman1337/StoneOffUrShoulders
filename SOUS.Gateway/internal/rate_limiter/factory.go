package rate_limiter

import (
	"fmt"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
)

func UseRateLimiter(cfg shared.RateLimiterConfig) gin.HandlerFunc {
	rl, err := getRateLimiter(cfg)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		fmt.Println("Before request")

		isAllowed := rl.IsAllowed(c.ClientIP())

		if !isAllowed {
			c.AbortWithStatus(429)
			return
		}

		c.Next()

		fmt.Println("After request")
	}
}

type RateLimiter interface {
	IsAllowed(key string) bool
}

var limiters = map[shared.RateLimiterStrategy]func(config shared.RateLimiterConfig) RateLimiter{
	"FixedWindow":   NewFixedWindowRateLimiter,
	"SlidingWindow": NewSlidingWindowRateLimiter,
	"TokenBucket":   NewTokenBucketRateLimiter,
}

func getRateLimiter(cfg shared.RateLimiterConfig) (RateLimiter, error) {
	strategy := cfg.Strategy

	result, ok := limiters[strategy]
	if !ok {
		return nil, fmt.Errorf("не найден лимитер соотвествующий ключу %s", strategy)
	} else {
		return result(cfg), nil
	}
}
