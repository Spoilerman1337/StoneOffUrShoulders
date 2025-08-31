package load_balancer

import (
	"fmt"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
)

type LoadBalancer interface {
	Next(c *gin.Context) *shared.Destination
}

var balancers = map[shared.LoadBalancerStrategy]func(cluster *shared.Cluster) LoadBalancer{
	"RoundRobin":            NewRoundRobinLoadBalancer,
	"WeightedRoundRobin":    NewWeightedRoundRobinLoadBalancer,
	"LeastRequests":         NewLeastRequestsBalancer,
	"WeightedLeastRequests": nil,
	"Random":                NewRandomLoadBalancer,
	"IPHash":                NewIPHashLoadBalancer,
}

func GetLoadBalancer(cluster *shared.Cluster) (LoadBalancer, error) {
	strategy := cluster.LoadBalancer

	result, ok := balancers[strategy]
	if !ok {
		return nil, fmt.Errorf("не найден балансировщик соотвествующий ключу %s", strategy)
	} else {
		return result(cluster), nil
	}
}
