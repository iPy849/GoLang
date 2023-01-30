package middlewares

import (
	"os"

	"fiber.example/handlers/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protects routes with JWT auth
func JWTProtected() func(*fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(utils.Response{Success: false, Message: "Missing or malformed JWT"})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(utils.Response{Success: false, Message: "Missing or malformed JWT"})
	}
}
