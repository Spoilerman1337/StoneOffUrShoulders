package config

type ServerConfig struct {
	Port string
}

type RouterConfig struct {
	Routes   []*Route   `yaml:"routes"`
	Clusters []*Cluster `yaml:"clusters"`
}

type Route struct {
	ClusterId string   `yaml:"clusterId"`
	Mask      string   `yaml:"mask"`
	Methods   []string `yaml:"methods"`
}

type Cluster struct {
	Name         string               `yaml:"name"`
	Destinations []*Destination       `yaml:"destinations"`
	LoadBalancer LoadBalancerStrategy `yaml:"loadBalancer"`
}

type Destination struct {
	Url string `yaml:"url"`
}

type LoadBalancerStrategy string

const (
	RoundRobin       LoadBalancerStrategy = "RoundRobin"
	LeastConnections LoadBalancerStrategy = "LeastConnections"
	Random           LoadBalancerStrategy = "Random"
)
