package shared

import "sync/atomic"

type Route struct {
	ClusterId string   `yaml:"clusterId"`
	Mask      string   `yaml:"mask"`
	Methods   []string `yaml:"methods"`
}

type Cluster struct {
	Destinations []*Destination       `yaml:"destinations"`
	LoadBalancer LoadBalancerStrategy `yaml:"loadBalancer"`
}

type Destination struct {
	Url               string       `yaml:"url"`
	Weight            int          `yaml:"weight"`
	activeConnections atomic.Int32 `yaml:"-"`
}

type RateLimiterConfig struct {
	Limit    int
	Strategy RateLimiterStrategy
	Rate     int64
}

func (d *Destination) ActiveConnections() int32 {
	return d.activeConnections.Load()
}

func (d *Destination) IncrementConnections() {
	d.activeConnections.Add(1)
}

func (d *Destination) DecrementConnections() {
	d.activeConnections.Add(-1)
}

type LoadBalancerStrategy string
type RateLimiterStrategy string

const (
	RoundRobin            LoadBalancerStrategy = "RoundRobin"
	WeightedRoundRobin    LoadBalancerStrategy = "WeightedRoundRobin"
	LeastRequests         LoadBalancerStrategy = "LeastRequests"
	WeightedLeastRequests LoadBalancerStrategy = "WeightedLeastRequests"
	IPHash                LoadBalancerStrategy = "IPHash"
	Random                LoadBalancerStrategy = "Random"
)

const (
	FixedWindow   RateLimiterStrategy = "FixedWindow"
	SlidingWindow RateLimiterStrategy = "SlidingWindow"
	TokenBucket   RateLimiterStrategy = "TokenBucket"
	LeakyBucket   RateLimiterStrategy = "LeakyBucket"
)
