package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/bxra2/7aweet/utils"
	"github.com/bxra2/7aweet/views"
	"github.com/gofiber/fiber/v2"
)

func (app *App) SearchTerm(c *fiber.Ctx) error {

	queryParam := c.Query("q")

	if queryParam == "" {
		return c.SendString("No query parameter found")
	}

	foundTerms, err := models.FindTermsByWord(app.DB, queryParam)
	if err != nil {
		log.Println("Error retrieving terms:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving terms")
	}
	return utils.Render(c, views.SearchPage(foundTerms, queryParam))
}
