package main

import (
	"gateway/internal/config"
	"gateway/internal/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitServerConfiguration()

	g := gin.Default()

	routing.InitRouting(g)

	err := g.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
