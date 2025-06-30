package routes

import (
	"be-ruang-warga/config" // Import package config
	ruangRiungDelivery "be-ruang-warga/internal/ruangriung/delivery"
	ruangRiungUsecase "be-ruang-warga/internal/ruangriung/usecase"
	userDelivery "be-ruang-warga/internal/user/delivery"
	userUsecase "be-ruang-warga/internal/user/usecase"

	"github.com/gin-gonic/gin"
	// firebaseAuth "firebase.google.com/go/v4/auth" // Tidak perlu import di sini jika diambil dari config.AuthClient
	// "gorm.io/gorm" // Tidak perlu import di sini jika diambil dari config.DB
)

// Tidak perlu menerima parameter db dan authClient lagi jika diambil dari config
func RegisterRoutes(router *gin.Engine) {
	db := config.DB           // Akses dari global variable
	authClient := config.AuthClient // Akses dari global variable
	api := router.Group("/api")

	rrUC := ruangRiungUsecase.NewRuangRiungUsecase(db)
	uUC := userUsecase.NewUserUsecase(db)

	// Lewatkan authClient ke NewUserHandler
	userDelivery.NewUserHandler(api, uUC, authClient) // <-- Perubahan di sini!
	ruangRiungDelivery.NewRuangRiungHandler(api, rrUC)
}