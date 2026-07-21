package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultCategories []string `json:"defaultCategories"`
	Projects          []string `json:"projects"`
	LastProject       int      `json:"lastProject"`
}

type ConfigStorage struct {
	configPath string
}

func NewConfigStorage(projectsDir string) *ConfigStorage {
	return &ConfigStorage{
		configPath: filepath.Join(projectsDir, "config.json"),
	}
}

func (s *ConfigStorage) Load() (*Config, error) {
	data, err := os.ReadFile(s.configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (s *ConfigStorage) Save(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.configPath, data, 0644)
}
