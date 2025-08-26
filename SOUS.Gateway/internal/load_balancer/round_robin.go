package load_balancer

import (
	"gateway/internal/shared"
	"sync/atomic"
)

type RoundRobinLoadBalancer struct {
	Destinations []*shared.Destination
	Counter      atomic.Int32
}

func (lb *RoundRobinLoadBalancer) Next() string {
	idx := lb.Counter.Add(1)

	return lb.Destinations[int(idx)%len(lb.Destinations)].Url
}

func NewRoundRobinLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &RoundRobinLoadBalancer{
		Destinations: cluster.Destinations,
		Counter:      atomic.Int32{},
	}
}
