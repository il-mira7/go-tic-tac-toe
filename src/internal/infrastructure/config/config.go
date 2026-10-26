package config

import (
	"log"
	"os"
)

type Config struct {
	DB  DatabaseConfig
	JWT JWTConfig
}

type DatabaseConfig struct {
	URL string
}

type JWTConfig struct {
	Secret string
}

func NewConfig() *Config {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("environment variable DATABASE_URL is required but not set")
	}
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("environment variable JWT_SECRET is required but not set")
	}

	return &Config{
		DB:  DatabaseConfig{URL: url},
		JWT: JWTConfig{Secret: secret},
	}
}
