package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

///	CONFIG

type Service struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
}

type Config struct {
	Host     string             `yaml:"host"`
	Port     string             `yaml:"port"`
	JWTKey   string             `yaml:"jwtKey"`
	Services map[string]Service `yaml:"services"`
}

func New(path string) (config *Config) {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error in open config file: %v", err)
	}
	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatalf("Error in marshal config file: %v", err)
	}
	log.Println("Config was got successfully")
	return
}
