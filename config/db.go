package config

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	domainRuangRiung "be-ruang-warga/internal/ruangriung/domain"
	domainUser "be-ruang-warga/internal/user/domain"
)

var (
	DB         *gorm.DB
	AuthClient *auth.Client
)

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	DB = database
	fmt.Println("Success to connect Database")

	err = DB.AutoMigrate(&domainUser.User{})
	DB.AutoMigrate(&domainRuangRiung.RuangRiung{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}
	fmt.Println("Database migration completed successfully!")
}

func InitFirebase() {
	ctx := context.Background()

	serviceAccountJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if serviceAccountJSON == "" {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS_JSON environment variable is not set. Please add your Firebase service account JSON content directly as an env variable in Railway.")
	}

	sa := option.WithCredentialsJSON([]byte(serviceAccountJSON))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	fmt.Println("Firebase Admin SDK initialized successfully.")

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Firebase Auth client: %v", err)
	}
	fmt.Println("Firebase Auth client obtained.")
}
