package config

import (
	"fmt"
	"load_balancer/pkg/logger"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Servers []ServerConfig `yaml:"servers"`
	Port    string         `yaml:"port"`
}

type ServerConfig struct {
	URL string `yaml:"url"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if len(data) == 0 {
		logger.Log.Error("Config file is empty", zap.String("path", path))
		return nil, fmt.Errorf("config file is empty")
	}
	if err != nil {
		logger.Log.Error("Failed to read config file", zap.Error(err))
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Log.Error("Failed to unmarshal config file", zap.Error(err))
		return nil, err
	}
	return &config, nil
}
