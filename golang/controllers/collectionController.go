package controllers

import (
	"log"

	"github.com/bxra2/7aweet/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllcollections handles the retrieval of all collections from the database
func (app *App) GetAllCollections(c *fiber.Ctx) error {
	// Get all collections from the model, passing the DB instance
	collections, err := models.GetAllCollections(app.DB)

	// Handle any error that occurs during the retrieval
	if err != nil {
		log.Println("Error retrieving collections:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving collections")
	}

	return c.JSON(collections)
}
