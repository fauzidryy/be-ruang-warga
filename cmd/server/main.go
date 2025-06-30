package main

import (
	"be-ruang-warga/config"
	"be-ruang-warga/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.InitFirebase() // <-- Panggil fungsi inisialisasi Firebase di sini!

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ruang Warga API is running!",
		})
	})

	// Sekarang, lewati config.DB dan config.AuthClient ke RegisterRoutes
	routes.RegisterRoutes(router) // <-- Parameter DB dan AuthClient akan diakses dari global var di package config
    // routes.RegisterRoutes(router, config.DB, config.AuthClient) // Jika kamu lebih suka passing explicitly

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("🚀 Server running on port", port)
	log.Fatal(router.Run(":" + port))
}