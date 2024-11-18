package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIToken string `yaml:"token"`
}

func GetConfiguration(configPath string, cfg interface{}) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return err
	}

	return nil
}
