package router

import (
	"fiber.example/handlers"
	"fiber.example/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Middlewares

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
