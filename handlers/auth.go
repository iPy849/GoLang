package handlers

import (
	"fmt"
	"os"
	"time"

	"fiber.example/handlers/utils"
	personModel "fiber.example/models/person"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthModel struct {
	Id        int64
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
}

// Handler para login en base a Nombres de Usuarios
// @Tags Auth
// @description Get auth JWT token if correct data
// @Param firstName formData string true "First Name"
// @Param lastName formData string true "Last Name"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500
// @Router /auth/login [post]
func LoginHandler(c *fiber.Ctx) error {
	// Recibe cuerpo del request y lo parsea
	userInfo := new(AuthModel)
	err := c.BodyParser(userInfo)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Busca usuario en BD
	person, err := personModel.RetrievePersonByNames(userInfo.FirstName, userInfo.LastName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	} else if person == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Crea el token y sus claims
	_token := jwt.New(jwt.SigningMethodHS256)
	claims := _token.Claims.(jwt.MapClaims)
	claims["username"] = fmt.Sprintf("%s %s", person.FirstName, person.LastName)
	claims["user_id"] = person.Id
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix() // Expiración

	// Genera token codificado
	secret := os.Getenv("JWT_SECRET_KEY")
	token, err := _token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.JSON(utils.Response{Success: true, Data: token})
}
