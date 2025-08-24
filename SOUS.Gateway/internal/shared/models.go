package shared

type Route struct {
	ClusterId string   `yaml:"clusterId"`
	Mask      string   `yaml:"mask"`
	Methods   []string `yaml:"methods"`
}

type Cluster struct {
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
