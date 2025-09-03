package config

import "gateway/internal/shared"

type ServerConfig struct {
	Port        string
	Https       bool
	RateLimiter RateLimiter
}

type RouterConfig struct {
	Routes   []*shared.Route            `yaml:"routes"`
	Clusters map[string]*shared.Cluster `yaml:"clusters"`
}

type RateLimiter struct {
	Limit    int
	Strategy RateLimiterStrategy
}

type RateLimiterStrategy string

const (
	FixedWindow   RateLimiterStrategy = "FixedWindow"
	SlidingWindow RateLimiterStrategy = "SlidingWindow"
	TokenBucket   RateLimiterStrategy = "TokenBucket"
	LeakyBucket   RateLimiterStrategy = "LeakyBucket"
)
