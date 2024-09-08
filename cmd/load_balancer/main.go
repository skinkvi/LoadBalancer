package main

import (
	"load_balancer/internal/config"
	"load_balancer/internal/server"
	"load_balancer/pkg/logger"
)

func main() {
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Log.Sugar().Fatalf("Failed to load config: %v", err)
	}

	server.StartServer(config)
}
