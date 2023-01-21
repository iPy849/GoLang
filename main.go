package main

import (
	"fiber.example/database"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/custom/:foo", func(c *fiber.Ctx) error {
		/*
			La ventaja de Fiber está en que trabaja con la menor cantidad de
			memoria posible, por eso se recomienda que uses las funciones y
			características del contexto (c) dentro del handler y que no mantengas
			los valores en memoria (pero si quieres hacer eso, cópialos y recuerda que
			tienes la función copy para slices).
		*/
		p := c.Params("foo")
		b := make([]byte, len(p))
		copy(b, p)
		return c.SendString(string(b))
	})

	app.Listen(":3000")
	// log.Fatal(app.Listen(":3000"))
}
