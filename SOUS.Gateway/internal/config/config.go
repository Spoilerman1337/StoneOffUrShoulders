package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitServerConfiguration() ServerConfig {
	viper.SetDefault("server.port", 8080)
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	errRead := viper.ReadInConfig()
	if errRead != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errRead))
	}

	var config ServerConfig
	errUnmarshal := viper.UnmarshalKey("server", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errUnmarshal))
	}

	return config
}

func InitRouterConfiguration() RouterConfig {
	viper.SetDefault("router.routes.methods", []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})
	viper.SetDefault("router.clusters.loadBalancer", "RoundRobin")
	viper.SetConfigName("router")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	errRead := viper.ReadInConfig()
	if errRead != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errRead))
	}

	var config RouterConfig
	errUnmarshal := viper.UnmarshalKey("router", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errUnmarshal))
	}

	return config
}
