package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configure()

	g := gin.Default()

	gwPort := ":" + viper.GetString("gateway.port")
	g.Run(gwPort)
}

func configure() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
