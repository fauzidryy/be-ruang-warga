package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️  No .env file found. Continuing...")
		} else {
			log.Println("✅ .env loaded locally")
		}
	} else {
		log.Println("🌐 Running on Railway, skipping .env load")
	}
}
