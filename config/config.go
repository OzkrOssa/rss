package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host []string `yaml:"host"`
}

func LoadConfig() ([]string, error) {

	yamlFile, _ := filepath.Abs("mikrotik.yaml")
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config.Host, nil
}
