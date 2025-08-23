package routing

import (
	"fmt"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func initProxy(g *gin.Engine, routes []*shared.Route, clusters []*shared.Cluster) {
	for _, route := range routes {
		cl, err := findCluster(clusters, route.ClusterId)
		if err != nil {
			panic(err)
		}

		targetUrl, _ := url.Parse(cl.Destinations[0].Url)
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)
		proxy.Director = func(req *http.Request) {
			transformedUrl, _ := strings.CutPrefix(req.URL.Path, "/api")

			req.URL.Scheme = targetUrl.Scheme
			req.URL.Host = targetUrl.Host
			req.URL.Path = transformedUrl
		}

		g.Any(route.Mask, func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}
}

func findCluster(clusters []*shared.Cluster, clusterId string) (*shared.Cluster, error) {
	for _, cluster := range clusters {
		if cluster.Name == clusterId {
			return cluster, nil
		}
	}

	return nil, fmt.Errorf("кластер с id %s не найден", clusterId)
}
