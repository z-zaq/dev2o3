package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}