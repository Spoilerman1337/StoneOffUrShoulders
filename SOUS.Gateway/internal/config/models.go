package config

import "gateway/internal/shared"

type ServerConfig struct {
	Port        string
	Https       bool
	RateLimiter shared.RateLimiterConfig
}

type RouterConfig struct {
	Routes   []*shared.Route            `yaml:"routes"`
	Clusters map[string]*shared.Cluster `yaml:"clusters"`
}
