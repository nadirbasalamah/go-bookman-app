package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	if os.Getenv("APP_MODE") == "prod" {
		return os.Getenv(key)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
