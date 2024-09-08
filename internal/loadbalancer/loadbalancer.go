package loadbalancer

import (
	config "load_balancer/internal/config"
	"load_balancer/pkg/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	URL         *url.URL
	connections uint64
}

type LoadBalancer struct {
	servers []*Server
	mu      sync.Mutex
}

func NewLoadBalancer(config *config.Config) *LoadBalancer {
	servers := make([]*Server, len(config.Servers))
	for i, serverConfig := range config.Servers {
		serverURL, err := url.Parse(serverConfig.URL)
		if err != nil {
			logger.Log.Error("Failed to parse server URL", zap.Error(err))
			return nil
		}
		servers[i] = &Server{URL: serverURL}
	}
	return &LoadBalancer{
		servers: servers,
	}
}

func (lb *LoadBalancer) NextServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var leastConnServer *Server
	leastConn := uint64(0)
	for _, server := range lb.servers {
		if server.connections < leastConn {
			leastConn = server.connections
			leastConnServer = server
		}
	}
	leastConnServer.connections++
	return leastConnServer
}

func (lb *LoadBalancer) DecreaseConnection(server *Server) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server.connections--
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := lb.NextServer()
	proxy := httputil.NewSingleHostReverseProxy(server.URL)
	proxy.Director = func(req *http.Request) {
		req.Host = server.URL.Host
		req.URL.Scheme = server.URL.Scheme
		req.URL.Host = server.URL.Host
	}

	proxy.ServeHTTP(w, r)
	lb.DecreaseConnection(server)
}
