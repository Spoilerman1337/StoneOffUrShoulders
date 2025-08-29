package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"sync"
)

type WeightedRoundRobinLoadBalancer struct {
	Destinations  []*shared.Destination
	Counter       int
	CurrentServed int
	Mutex         sync.Mutex
}

func (lb *WeightedRoundRobinLoadBalancer) Next(c *gin.Context) string {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	idx := lb.Counter
	dest := lb.Destinations[idx%(len(lb.Destinations))]

	if lb.CurrentServed < dest.Weight {
		lb.CurrentServed += 1
	} else {
		lb.CurrentServed = 0
		lb.Counter += 1
		idx = lb.Counter
	}

	return lb.Destinations[idx%len(lb.Destinations)].Url
}

func NewWeightedRoundRobinLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &WeightedRoundRobinLoadBalancer{
		Destinations:  cluster.Destinations,
		Counter:       0,
		CurrentServed: 0,
		Mutex:         sync.Mutex{},
	}
}
