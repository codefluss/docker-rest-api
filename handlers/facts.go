package handlers

import (
	"github.com/codefluss/docker-rest-api/database"
	"github.com/codefluss/docker-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

 func ListFacts(c *fiber.Ctx) error {
 	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(fiber.StatusOK).JSON(facts)
}


 func CreateFact(c *fiber.Ctx) error {
 	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "An error occurred while parsing the request body",
		})
	}

	database.DB.Db.Create(&fact)	
	return c.Status(fiber.StatusCreated).JSON(fact)
 }