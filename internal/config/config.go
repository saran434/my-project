package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func Load() DBConfig {
	var cfg DBConfig
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("failed to read file %s: %v", "config.yml", err)
	}
	fmt.Println("config.yml read successfully")
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to unmarshal YAML: %v", err)
	}
	fmt.Println("config.yml unmarshaled successfully")
	return cfg
}

func Pick(cfgVal, envKey string) string {
	if cfgVal != "" {
		return cfgVal
	}
	return os.Getenv(envKey)
}

func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretjwtkey"
	}
	return secret
}
