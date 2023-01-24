package handlers

import (
	"fmt"
	"strconv"

	personModel "fiber.example/models/person"
	"github.com/gofiber/fiber/v2"
)

// Handler que recupera varias personas. Acepta queries "limit" y "offset"
func GetManyPersonHandler(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	data, err := personModel.RetrieveNPeople(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if len(data) == 0 {
		return c.Status(fiber.StatusNoContent).SendString("No content")
	} else {
		return c.JSON(data)
	}
}

// Handler que recupera una persona según un parámetro "id" en la url
func GetPersonHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("A")
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	person, err := personModel.RetrievePerson(id)
	if err != nil {
		fmt.Println("B")
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if person == nil {
		return c.Status(fiber.StatusNoContent).SendString("No content")
	}
	return c.JSON(person)
}

// Handler que crea una persona. La información viene codificada como form-data
func CreatePersonHandler(c *fiber.Ctx) error {
	// NOTE: Parsea la entrada a una estructura
	person := new(personModel.Person)
	c.BodyParser(person)

	id, err := personModel.CreatePerson(person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// Handler que actualiza a una persona en base a información codificada en form-data y un id de la url
func UpdatePersonHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	person := new(personModel.Person)
	c.BodyParser(person)

	err = personModel.UpdatePerson(id, person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return GetPersonHandler(c)
}

// Handler que elimina una persona.
func DeletePersonHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	err = personModel.DeletePerson(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(fiber.StatusNoContent).SendString("No content")

}
