package router

import (
	_ "fiber.example/docs"
	"fiber.example/handlers"
	"fiber.example/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// Static Assets
	assetsGroup := app.Group("/static")
	assetsGroup.Static("/img", "./assets/img")

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New()) // Recuperación cuando crashea el programa

	// Default index, testing purposes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola amigo")
	})

	// Documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handlers.LoginHandler)

	// Person CRUD
	// Uso el middleware de JWT para todo el grupo de rutas /person
	person := app.Group("/person", middlewares.JWTProtected())
	person.Get("", handlers.GetManyPersonHandler)
	person.Get("/:id<int>", handlers.GetPersonHandler)
	person.Post("", handlers.CreatePersonHandler)
	person.Put("/:id<int>", handlers.UpdatePersonHandler)
	person.Delete("/:id<int>", handlers.DeletePersonHandler)
}
