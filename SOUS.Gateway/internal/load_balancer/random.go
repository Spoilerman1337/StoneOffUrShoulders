package load_balancer

import (
	"gateway/internal/shared"
	"math/rand"
)

type RandomLoadBalancer struct {
	Destinations []*shared.Destination
}

func (lb RandomLoadBalancer) Next() string {
	randomizationDegree := len(lb.Destinations)

	return lb.Destinations[rand.Intn(randomizationDegree)].Url
}

func NewRandomLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &RandomLoadBalancer{
		Destinations: cluster.Destinations,
	}
}
