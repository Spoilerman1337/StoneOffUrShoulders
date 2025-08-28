package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type RandomLoadBalancer struct {
	Destinations []*shared.Destination
}

func (lb RandomLoadBalancer) Next(c *gin.Context) string {
	randomizationDegree := len(lb.Destinations)

	return lb.Destinations[rand.Intn(randomizationDegree)].Url
}

func NewRandomLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &RandomLoadBalancer{
		Destinations: cluster.Destinations,
	}
}
