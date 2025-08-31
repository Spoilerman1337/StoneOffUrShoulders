package load_balancer

import (
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"math"
	"sync"
)

type WeightedLeastRequestsBalancer struct {
	Destinations []*shared.Destination
	Mutex        sync.Mutex
}

func (lb *WeightedLeastRequestsBalancer) Next(c *gin.Context) *shared.Destination {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	minScore := float32(math.MaxFloat32)
	var chosenDestination *shared.Destination

	for _, dest := range lb.Destinations {
		score := float32(dest.ActiveConnections()) / float32(dest.Weight)
		if minScore > score {
			minScore = score
			chosenDestination = dest
		}
	}

	if chosenDestination != nil {
		chosenDestination.IncrementConnections()
	}

	return chosenDestination
}

func NewWeightedLeastRequestsBalancer(cluster *shared.Cluster) LoadBalancer {
	return &WeightedLeastRequestsBalancer{
		Destinations: cluster.Destinations,
		Mutex:        sync.Mutex{},
	}
}
