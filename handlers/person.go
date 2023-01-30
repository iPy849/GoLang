package handlers

import (
	"strconv"

	"fiber.example/handlers/utils"
	personModel "fiber.example/models/person"
	"github.com/gofiber/fiber/v2"
)

// Handler que recupera varias personas.
// @Tags Person
// @description Get many users
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} utils.Response
// @Success 204 {object} utils.Response
// @Failure 400
// @Failure 500
// @Router /person [get]
// @Security JWT
func GetManyPersonHandler(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	data, err := personModel.RetrieveNPeople(limit, offset)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(data) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(utils.Response{Success: true, Data: nil})
	} else {
		return c.JSON(utils.Response{Success: true, Data: data})
	}
}

// Handler que recupera una persona según un parámetro "id" en la url
// @Tags Person
// @description Gets a user based in it's user id
// @Param id path int true "Id"
// @Success 200 {object} personModel.Person
// @Success 204
// @Failure 400
// @Failure 500
// @Router /person/{id} [get]
// @Security JWT
func GetPersonHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	person, err := personModel.RetrievePerson(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if person == nil {
		return c.Status(fiber.StatusNoContent).SendString("No content")
	}
	return c.JSON(person)
}

// Handler que crea una persona. La información viene codificada como form-data
// @Tags Person
// @description Creates a person based in form-data information
// @Param firstName formData string true "First Name"
// @Param lastName formData string true "Last Name"
// @Param email formData string true "Email"
// @Param gender formData string true "Gender"
// @Param dateOfBirth formData string true "Date Of Birth"
// @Param countryOfBirth formData string true "Country Of Birth"
// @Success 200 {object} utils.Response
// @Success 204
// @Failure 400
// @Failure 500
// @Router /person [post]
// @Security JWT
func CreatePersonHandler(c *fiber.Ctx) error {
	// NOTE: Parsea la entrada a una estructura
	person := new(personModel.Person)
	c.BodyParser(person)

	id, err := personModel.CreatePerson(person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(utils.Response{Success: true, Data: fiber.Map{"id": id}})
}

// Handler que actualiza a una persona en base a información codificada en form-data y un id de la url
// @Tags Person
// @description Updates a person based in form-data information
// @Param id path int true "Id"
// @Param firstName formData string false "First Name"
// @Param lastName formData string false "Last Name"
// @Param email formData string false "Email"
// @Param gender formData string false "Gender"
// @Param dateOfBirth formData string false "Date Of Birth"
// @Param countryOfBirth formData string false "Country Of Birth"
// @Success 200 {object} personModel.Person
// @Success 204
// @Failure 400
// @Failure 500
// @Router /person/{id} [put]
// @Security JWT
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
// @Tags Person
// @description Deletes a person based in it's id
// @Param id path int true "Id"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /person/{id} [delete]
// @Security JWT
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
