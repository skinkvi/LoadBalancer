package server

import (
	"fmt"
	"load_balancer/internal/config"
	"load_balancer/internal/loadbalancer"
	"load_balancer/pkg/logger"
	"net/http"

	"go.uber.org/zap"
)

func StartServer(config *config.Config) {
	lb := loadbalancer.NewLoadBalancer(config)

	http.Handle("/", lb)

	logger.Log.Info(fmt.Sprintf("Starting server on port %s", config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)
	if err != nil {
		logger.Log.Error("Failed to start server", zap.Error(err))
	}
}
