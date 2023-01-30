package main

import (
	"fiber.example/database"
	"fiber.example/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title Fiber Example API
// @version 0.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email frank.ortegaca@htech.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		message := "Error loading .env file"
		panic(message)
	}

	// Database connection
	database.NewDbConnectionPool()
	defer database.DbConnectionPool.Close()

	// Fiber Setup
	app := fiber.New()

	// app.Get("/custom/:foo", func(c *fiber.Ctx) error {
	// 	/*
	// 		La ventaja de Fiber está en que trabaja con la menor cantidad de
	// 		memoria posible, por eso se recomienda que uses las funciones y
	// 		características del contexto (c) dentro del handler y que no mantengas
	// 		los valores en memoria (pero si quieres hacer eso, cópialos y recuerda que
	// 		tienes la función copy para slices).
	// 	*/
	// 	p := c.Params("foo")
	// 	b := make([]byte, len(p))
	// 	copy(b, p)
	// 	return c.SendString(string(b))
	// })

	// Setup routes
	router.SetupRoutes(app)

	app.Listen(":3000")
	// log.Fatal(app.Listen(":3000"))
}
