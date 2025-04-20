package controllers

import (
	"log"
	"strconv"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
)

func (app *App) SearchTerm(c *fiber.Ctx) error {

	queryParam := c.Query("q")
	includeDesc := c.Query("description", "0")
	includeFr := c.Query("french", "0")
	includeDe := c.Query("german", "0")
	/// Pagination Params
	pageStr := c.Query("page", "1")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil || pageNum < 1 {
		pageNum = 1 // Default to page 1 if conversion fails or invalid
	}

	// Set a default for the limit parameter
	limitStr := c.Query("limit", "10")
	limitNum, err := strconv.Atoi(limitStr)
	if err != nil || limitNum < 1 {
		limitNum = 10 // Default to limit of 10
	}

	includeDescBool := includeDesc == "1"
	includeFrBool := includeFr == "1"
	includeDeBool := includeDe == "1"

	if queryParam == "" {
		queryParam = "undefined"
	}

	termCount, foundTerms, sources, domains, err := models.FindTermsByWord(app.DB, queryParam, includeDescBool, includeFrBool, includeDeBool, pageNum, limitNum)
	if err != nil {
		log.Println("Error retrieving terms:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving terms")
	}
	return c.JSON(fiber.Map{"Count": termCount, "sources": sources, "domains": domains, "terms": foundTerms})
}

func (app *App) Find10RandomWords(c *fiber.Ctx) error {
	foundTerms, err := models.GetRandomTerms(app.DB)
	if err != nil {
		log.Println("Error retrieving terms:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving terms")
	}
	return c.JSON(foundTerms)
}
