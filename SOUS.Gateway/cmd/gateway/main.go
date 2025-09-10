package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/rate_limiter"
	"gateway/internal/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitServerConfiguration()

	g := gin.Default()

	routing.InitRouting(g)

	stpErr := g.SetTrustedProxies([]string{"127.0.0.1"})
	if stpErr != nil {
		fmt.Println(fmt.Errorf("Error setting trusted proxies: %v", stpErr))
	}
	g.Use(rate_limiter.UseRateLimiter(cfg.RateLimiter))

	err := g.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
