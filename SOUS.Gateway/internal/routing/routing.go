package routing

import (
	"gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func InitRouting(g *gin.Engine) {
	cfg := config.InitRouterConfiguration()

	ensureRouteUniqueness(cfg.Routes)
	ensureClusterUniqueness(cfg.Clusters)

	initProxy(g, cfg.Routes, cfg.Clusters)
}
