package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GRPC GRPCConf
}

type GRPCConf struct {
	Host string
	Port string
}

func NewConfig() Config {
	return Config{}
}

func LoadConfig(path string) (*Config, error) {
	result, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't parse config: %w", err)
	}

	config := NewConfig()
	err = yaml.Unmarshal(result, &config)
	if err != nil {
		return nil, fmt.Errorf("cant unmarshal config: %w", err)
	}

	return &config, nil
}
