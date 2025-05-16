package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackConfig struct {
	Packs []int `json:"packs"`
}

func LoadPackConfig(path string) (*PackConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg PackConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}
