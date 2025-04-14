package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

// GetAllSources handles the retrieval of all sources from the database
func (app *App) GetAllSources(c *fiber.Ctx) error {
	// Get all sources from the model, passing the DB instance
	sources, err := models.GetAllSources(app.DB)

	// Handle any error that occurs during the retrieval
	if err != nil {
		log.Println("Error retrieving sources:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving sources")
	}

	return c.JSON(sources)
}
