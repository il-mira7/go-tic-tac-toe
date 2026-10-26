package di

import (
	"log"
	"os"
)

func NewDatabaseConfig() string {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("environment variable DATABASE_URL is required but not set")
	}

	return url
}
