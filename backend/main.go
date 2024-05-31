package main

import "github.com/gofiber/fiber/v2"

func main() {
	// creating app
	app := fiber.New()

	// group routes to v1
	apiv1 := app.Group("/api/v1")

	// routes
	app.Get("/", handleHome)
	apiv1.Get("/", handleIndex)

	// server listing
	app.Listen(":3000")
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "server is running!"})
}

func handleIndex(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "welcome to hotel reservation index!"})
}
