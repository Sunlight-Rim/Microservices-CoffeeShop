package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

///	CONFIG

type Service struct {
	URL  string `yaml:"url"`
	Port int16  `yaml:"port"`
}

type Config struct {
	Host     string             `yaml:"host"`
	Port     string             `yaml:"port"`
	Services map[string]Service `yaml:"services"`
}

func New() (cfg *Config) {
	f, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error in open config file: %v", err)
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatalf("Error in marshal config file: %v", err)
	}
	return
}
