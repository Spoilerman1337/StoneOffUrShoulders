package routing

import (
	"fmt"
	"gateway/internal/shared"
)

func ensureClusterUniqueness(clusters []*shared.Cluster) {
	dict := make(map[string]*shared.Cluster)
	for _, cluster := range clusters {
		value, ok := dict[cluster.Name]
		if !ok {
			dict[cluster.Name] = cluster
		} else {
			panic(fmt.Errorf("обнаружен дубликат кластера: %q", value.Name))
		}
	}
}

func ensureRouteUniqueness(routes []*shared.Route) {
	dict := make(map[string]*shared.Route)
	for _, routes := range routes {
		value, ok := dict[routes.Mask]
		if !ok {
			dict[routes.Mask] = routes
		} else {
			panic(fmt.Errorf("обнаружен дубликат пути: %q", value.Mask))
		}
	}
}
