package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", handleHome)
	app.Listen(":3000")
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "server is running!"})
}
