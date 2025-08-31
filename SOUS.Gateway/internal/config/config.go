package config

import (
	"fmt"
	"gateway/internal/shared"
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
	setRoutesDefaults(config.Routes)
	setClustersDefaults(config.Clusters)

	return config
}

func setRoutesDefaults(routes []*shared.Route) {
	for _, route := range routes {
		if len(route.Methods) == 0 {
			route.Methods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"}
		}
	}
}

func setClustersDefaults(clusters map[string]*shared.Cluster) {
	for _, cluster := range clusters {
		if len(cluster.LoadBalancer) == 0 {
			cluster.LoadBalancer = "RoundRobin"
		}
		setDestinationDefaults(cluster.Destinations)
	}
}

func setDestinationDefaults(destinations []*shared.Destination) {
	for _, destination := range destinations {
		if destination.Weight == 0 {
			destination.Weight = 1
		}
	}
}
