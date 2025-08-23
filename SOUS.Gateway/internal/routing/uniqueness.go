package routing

import (
	"fmt"
	"gateway/internal/config"
)

func ensureClusterUniqueness(clusters []*config.Cluster) {
	dict := make(map[string]*config.Cluster)
	for _, cluster := range clusters {
		value, ok := dict[cluster.Name]
		if !ok {
			dict[cluster.Name] = cluster
		} else {
			panic(fmt.Errorf("обнаружен дубликат кластера: %q", value.Name))
		}
	}
}

func ensureRouteUniqueness(routes []*config.Route) {
	dict := make(map[string]*config.Route)
	for _, routes := range routes {
		value, ok := dict[routes.Mask]
		if !ok {
			dict[routes.Mask] = routes
		} else {
			panic(fmt.Errorf("обнаружен дубликат пути: %q", value.Mask))
		}
	}
}
