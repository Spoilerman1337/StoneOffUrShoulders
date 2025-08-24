package routing

import (
	"gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func InitRouting(g *gin.Engine) {
	cfg := config.InitRouterConfiguration()

	initProxy(g, cfg.Routes, cfg.Clusters)
}
