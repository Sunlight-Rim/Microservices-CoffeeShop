package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

///	CONFIG

type Service struct {
	URL string `yaml:"url"`
}

type Config struct {
	Host     string             `yaml:"host"`
	Port     string             `yaml:"port"`
	Services map[string]Service `yaml:"services"`
}

func New() (cfg *Config) {
	f, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		panic(err)
	}
	return
}

func GetSocket() string {
	cfg := New()
	return cfg.Host + ":" + cfg.Port
}
