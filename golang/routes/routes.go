package routes

import (
	"strings"

	"github.com/bxra2/7aweet/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, controller *controllers.App) {
	api := app.Group("/api")
	// set routes for api

	api.Get("/about", controller.LoadAboutPage)
	api.Get("/sources", controller.GetAllSources)
	api.Get("/collections", controller.GetAllCollections)
	api.Get("/terms/random", controller.Find10RandomWords)
	api.Get("/search", controller.SearchTerm)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	// set not found handler
	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	})
}

// set single page application fallback
func SPAFallback(indexPath string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api/") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
		}
		return c.SendFile(indexPath)
	}
}
