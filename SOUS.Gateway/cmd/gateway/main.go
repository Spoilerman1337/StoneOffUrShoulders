package main

import (
	"gateway/internal/config"
	"gateway/internal/load_balancer"
	"gateway/internal/routing"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"sync"
	"testing"
)

func main() {
	cfg := config.InitServerConfiguration()

	g := gin.Default()

	routing.InitRouting(g)

	TestRoundRobinRace(&testing.T{})

	err := g.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}

func TestRoundRobinRace(t *testing.T) {
	cluster := &shared.Cluster{
		Destinations: []*shared.Destination{
			{Url: "http://a"},
			{Url: "http://b"},
			{Url: "http://c"},
		},
	}
	lb := load_balancer.NewRoundRobinLoadBalancer(cluster)

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			_ = lb.Next()
			wg.Done()
		}()
	}
	wg.Wait()
}
