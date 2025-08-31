package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"math"
)

type LeastRequestsBalancer struct {
	Destinations []*shared.Destination
}

func (lb *LeastRequestsBalancer) Next(c *gin.Context) *shared.Destination {
	minActiveRequestPool := int32(math.MaxInt32)
	var chosenDestination *shared.Destination
	for _, dest := range lb.Destinations {
		if minActiveRequestPool > dest.ActiveConnections() {
			minActiveRequestPool = dest.ActiveConnections()
			chosenDestination = dest
		}
	}

	if chosenDestination != nil {
		chosenDestination.IncrementConnections()
	}

	return chosenDestination
}

func NewLeastRequestsBalancer(cluster *shared.Cluster) LoadBalancer {
	return &LeastRequestsBalancer{
		Destinations: cluster.Destinations,
	}
}
