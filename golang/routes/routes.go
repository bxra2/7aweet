package routes

import (
	"github.com/bxra2/7aweet/controllers"
	"github.com/bxra2/7aweet/utils"
	"github.com/bxra2/7aweet/views"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, controller *controllers.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return utils.Render(c, views.Home())
	// })

	app.Get("/api/about", controller.LoadAboutPage)

	app.Get("/api/sources", controller.GetAllSources)

	// app.Get("/api/suggestions", func(c *fiber.Ctx) error {
	// 	return utils.Render(c, views.Suggestions())
	// })
	app.Get("/api/search", controller.SearchTerm)

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return utils.Render(c, views.NotFound())
}
