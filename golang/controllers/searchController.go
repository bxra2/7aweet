package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
)

func (app *App) SearchTerm(c *fiber.Ctx) error {

	queryParam := c.Query("q")
	includeDesc := c.Query("description", "0")
	includeFr := c.Query("french", "0")
	includeDe := c.Query("german", "0")

	includeDescBool := includeDesc == "1"
	includeFrBool := includeFr == "1"
	includeDeBool := includeDe == "1"

	if queryParam == "" {
		queryParam = "undefined"
	}

	foundTerms, err := models.FindTermsByWord(app.DB, queryParam, includeDescBool, includeFrBool, includeDeBool)
	if err != nil {
		log.Println("Error retrieving terms:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving terms")
	}
	return c.JSON(foundTerms)
}
