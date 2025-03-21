package controllers

import (
	"log"
	"strconv"

	"github.com/bxra2/7aweet/models"
	"github.com/bxra2/7aweet/utils"
	"github.com/bxra2/7aweet/views"
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
	// Respond with JSON (can be HTML later if needed)
	return utils.Render(c, views.Sources(sources, strconv.Itoa(len(sources))))
}
