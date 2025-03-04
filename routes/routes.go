package routes

import (
	"github.com/a-h/templ"
	"github.com/bxra2/7aweet/views"
	"github.com/gofiber/fiber/v2"
)

func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return render(c, views.About())
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		queryParam := c.Query("q")

		if queryParam == "" {
			return c.SendString("No query parameter found")
		}
		return c.SendString("Query parameter 'q' value: " + queryParam)
	})

}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return render(c, views.NotFound())
}
