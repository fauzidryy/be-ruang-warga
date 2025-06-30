package config

import (
	"context" // Import ini
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"         // Import ini
	"firebase.google.com/go/v4/auth"              // Import ini
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"google.golang.org/api/option" // Import ini
)

var (
	DB         *gorm.DB
	AuthClient *auth.Client // <-- Tambahkan ini untuk menyimpan Firebase Auth Client
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

	// PENTING: Lakukan AutoMigrate di sini setelah DB terkoneksi
	// Contoh (jika kamu punya model User dan AdminRequest di be-ruang-warga/internal/user/domain):
	// import "be-ruang-warga/internal/user/domain"
	// err = DB.AutoMigrate(&domain.User{}, &domain.AdminRequest{})
	// if err != nil {
	//     log.Fatalf("Failed to auto-migrate database schema: %v", err)
	// }
	// fmt.Println("Database migration completed successfully!")
}

// --- FUNGSI BARU UNTUK INISIALISASI FIREBASE ---
func InitFirebase() {
	ctx := context.Background()

	serviceAccountKeyPath := os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_PATH")
	if serviceAccountKeyPath == "" {
		log.Fatal("FIREBASE_SERVICE_ACCOUNT_KEY_PATH environment variable is not set. Please specify the path to your Firebase service account key JSON file.")
	}

	sa := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	fmt.Println("Firebase Admin SDK initialized successfully.")

	AuthClient, err = app.Auth(ctx) // Assign ke global variable AuthClient
	if err != nil {
		log.Fatalf("Error getting Firebase Auth client: %v", err)
	}
	fmt.Println("Firebase Auth client obtained.")
}
// --- End of Firebase Initialization ---