package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
)

func (app *App) LoadAboutPage(c *fiber.Ctx) error {
	domains, err := models.GetDomainsWithTermCountsRaw(app.DB)
	if err != nil {
		log.Println("Error retrieving domains:", err)
		return c.Status(fiber.StatusInternalServerError).
			SendString("Error retrieving domains")
	}
	return c.JSON(domains)
}
