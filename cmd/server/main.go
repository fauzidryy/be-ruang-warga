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

	router := gin.Default()

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("🚀 Server running on port", port)
	log.Fatal(router.Run(":" + port))
}
