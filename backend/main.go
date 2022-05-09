package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) {
		c.Send("go fiber is health")
	})

	app.Listen(3000)
}
