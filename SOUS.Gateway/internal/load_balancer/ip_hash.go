package load_balancer

import (
	"fmt"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"hash/fnv"
)

type IPHashLoadBalancer struct {
	Destinations []*shared.Destination
}

func (lb *IPHashLoadBalancer) Next(c *gin.Context) string {
	hash := fnv.New32a()
	hashLen, _ := hash.Write([]byte(c.ClientIP()))
	var val int
	if hashLen > 0 {
		val = int(hash.Sum32())
	}

	fmt.Printf("Hash: %d. IP: %s", val, c.ClientIP())

	return lb.Destinations[val%len(lb.Destinations)].Url
}

func NewIPHashLoadBalancer(cluster *shared.Cluster) LoadBalancer {
	return &IPHashLoadBalancer{
		Destinations: cluster.Destinations,
	}
}
