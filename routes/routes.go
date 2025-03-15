package routes

import (
	"github.com/bxra2/7aweet/controllers"
	"github.com/bxra2/7aweet/utils"
	"github.com/bxra2/7aweet/views"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, controller *controllers.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, views.Home())
	})

	app.Get("/about", controller.LoadAboutPage)

	app.Get("/sources", controller.GetAllSources)

	app.Get("/suggestions", func(c *fiber.Ctx) error {
		return utils.Render(c, views.Suggestions())
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		queryParam := c.Query("q")

		if queryParam == "" {
			return c.SendString("No query parameter found")
		}
		return utils.Render(c, views.SearchPage(queryParam))
		// return c.SendString("Query parameter 'q' value: " + queryParam)
	})

}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return utils.Render(c, views.NotFound())
}
