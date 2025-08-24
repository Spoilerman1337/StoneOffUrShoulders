package config

import "gateway/internal/shared"

type ServerConfig struct {
	Port string
}

type RouterConfig struct {
	Routes   []*shared.Route            `yaml:"routes"`
	Clusters map[string]*shared.Cluster `yaml:"clusters"`
}
