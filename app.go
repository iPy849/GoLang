package main

import (
	"fiber_sockets/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	router.RoutesDefinition(app)

	log.Println(app.Listen(":3000"))

}
