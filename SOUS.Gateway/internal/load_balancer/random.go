package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type RandomLoadBalancer struct {
	Destinations []*shared.Destination
}

func (lb RandomLoadBalancer) Next(c *gin.Context) *shared.Destination {
	randomizationDegree := len(lb.Destinations)

	return lb.Destinations[rand.Intn(randomizationDegree)]
}

func NewRandomLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &RandomLoadBalancer{
		Destinations: cluster.Destinations,
	}
}
