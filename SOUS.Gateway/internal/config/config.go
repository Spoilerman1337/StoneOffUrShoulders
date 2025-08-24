package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitServerConfiguration() ServerConfig {
	v := viper.New()

	v.SetDefault("server.port", 8080)
	v.SetConfigName("server")
	v.SetConfigType("json")
	v.AddConfigPath("configs")
	errRead := v.ReadInConfig()
	if errRead != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errRead))
	}

	var config ServerConfig
	errUnmarshal := v.UnmarshalKey("server", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errUnmarshal))
	}

	return config
}

func InitRouterConfiguration() RouterConfig {
	v := viper.New()

	v.SetConfigName("router")
	v.SetConfigType("json")
	v.AddConfigPath("configs")
	errRead := v.ReadInConfig()
	if errRead != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errRead))
	}

	var config RouterConfig
	errUnmarshal := v.UnmarshalKey("router", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка инициализации: %w", errUnmarshal))
	}

	// К сожалению, viper не может установить конфиг по умолчанию внутри массива,
	// поэтому мы вынуждены доинициализировать вручную
	for _, route := range config.Routes {
		if len(route.Methods) == 0 {
			route.Methods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"}
		}
	}
	for _, cluster := range config.Clusters {
		if len(cluster.LoadBalancer) == 0 {
			cluster.LoadBalancer = "RoundRobin"
		}
	}

	return config
}
