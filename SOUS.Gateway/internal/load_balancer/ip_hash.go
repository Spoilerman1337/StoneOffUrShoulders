package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"hash/fnv"
)

type IPHashLoadBalancer struct {
	Destinations []*shared.Destination
}

func (lb *IPHashLoadBalancer) Next(c *gin.Context) *shared.Destination {
	hash := fnv.New32a()
	hashLen, _ := hash.Write([]byte(c.ClientIP()))
	var val int
	if hashLen > 0 {
		val = int(hash.Sum32())
	}

	return lb.Destinations[val%len(lb.Destinations)]
}

func NewIPHashLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &IPHashLoadBalancer{
		Destinations: cluster.Destinations,
	}
}
