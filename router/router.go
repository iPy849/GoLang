package router

import (
	"os"

	chat "fiber_sockets/websockets"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func RoutesDefinition(app *fiber.App) {

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New())

	// Site root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte(os.Getenv("WELCOME_MESSAGE")))
	})

	// WebSockets
	chatGroup := app.Group("ws")
	chatGroup.Get("/chat", websocket.New(chat.ChatHandler))

	// Handling 404
	app.Use(func(c *fiber.Ctx) error {
		url := c.Hostname() + c.Path()
		return c.Status(fiber.StatusNotFound).SendString("Sorry we can't find: " + url)
	})

}
