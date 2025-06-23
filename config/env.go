package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Println("⚠️  Failed to load .env file:", err)
		} else {
			log.Println("✅ .env file loaded")
		}
	} else {
		log.Println("🌐 No .env file found, assume running in production")
	}
}
