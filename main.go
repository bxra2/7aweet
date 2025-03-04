package main

import (
	"log"
	"os"

	"github.com/bxra2/7aweet/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("cannot find environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}
	app := fiber.New()

	routes.SetRoutes(app)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
