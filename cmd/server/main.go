package main

import (
	"be-ruang-warga/config"
	"be-ruang-warga/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	port := ":5434"
	fmt.Println("running in" + port)
	app.Listen(port)
}