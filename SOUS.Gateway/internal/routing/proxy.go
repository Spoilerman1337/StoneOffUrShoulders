package routing

import (
	"fmt"
	"gateway/internal/load_balancer"
	"gateway/internal/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func initProxy(g *gin.Engine, routes []*shared.Route, clusters map[string]*shared.Cluster) {
	for _, route := range routes {
		cl, ok := clusters[route.ClusterId]
		if !ok {
			panic("cluster id not exist")
		}

		targetUrl, _ := url.Parse(cl.Destinations[0].Url)
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)
		proxy.Director = func(req *http.Request) {
			transformedUrl, _ := strings.CutPrefix(req.URL.Path, "/api")

			req.URL.Scheme = targetUrl.Scheme
			req.URL.Host = targetUrl.Host
			req.URL.Path = transformedUrl
		}

		lb, err := load_balancer.GetLoadBalancer(cl)
		if err != nil {
			panic(err)
		}

		for _, method := range route.Methods {
			g.Handle(method, route.Mask, func(c *gin.Context) {
				newUrl := lb.Next()
				fmt.Println(fmt.Sprintf("Loadbalancer chose url: %s", newUrl))

				proxy.ServeHTTP(c.Writer, c.Request)
			})
		}
	}
}
