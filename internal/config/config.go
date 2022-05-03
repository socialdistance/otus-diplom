package config

import (
	"fmt"
	yaml3 "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Stats StatsConfig
}

type StatsConfig struct {
	LoadAvg bool
	CPU     bool
	Disk    bool
	Memory  bool
	// NetTop  bool
	// NetStat bool
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
	err = yaml3.Unmarshal(result, &config)
	if err != nil {
		return nil, fmt.Errorf("cant unmarshal config: %w", err)
	}

	return &config, nil
}
