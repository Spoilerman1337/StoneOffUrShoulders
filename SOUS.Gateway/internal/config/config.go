package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitServerConfiguration() ServerConfig {
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	errRead := viper.ReadInConfig()
	if errRead != nil {
		panic(errRead)
	}

	var config ServerConfig
	errUnmarshal := viper.UnmarshalKey("server", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка десериализации: %w", errUnmarshal))
	}

	return config
}

func InitRouterConfiguration() RouterConfig {
	viper.SetConfigName("router")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	errRead := viper.ReadInConfig()
	if errRead != nil {
		panic(errRead)
	}

	var config RouterConfig
	errUnmarshal := viper.UnmarshalKey("router", &config)
	if errUnmarshal != nil {
		panic(fmt.Errorf("ошибка десериализации: %w", errUnmarshal))
	}

	return config
}
