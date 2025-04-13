package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
)

func (app *App) GetAllDomains(c *fiber.Ctx) ([]models.Domain, error) {
	domains, err := models.GetAllDomains(app.DB)
	if err != nil {
		log.Println("Error retrieving sources:", err)
		c.Status(fiber.StatusInternalServerError).SendString("Error retrieving sources")
		return nil, err
	}
	return domains, nil
}
