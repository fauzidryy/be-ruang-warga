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

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ruang Warga API is running!",
		})
	})

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("🚀 Server running on port", port)
	log.Fatal(router.Run(":" + port))
}
