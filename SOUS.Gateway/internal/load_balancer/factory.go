package load_balancer

import (
	"fmt"
	"gateway/internal/shared"
)

type LoadBalancer interface {
	Next() string
}

var balancers = map[shared.LoadBalancerStrategy]func(cluster *shared.Cluster) LoadBalancer{
	"RoundRobin":       NewRoundRobinLoadBalancer,
	"LeastConnections": nil,
	"Random":           NewRandomLoadBalancer,
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
