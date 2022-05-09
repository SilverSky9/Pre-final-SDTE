package controllers

import (
	"sdte-backend/services"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		msg, _ := services.GetHealth()
		return c.Status(200).JSON(msg)
	})
	// app.Get("/person", func(c *fiber.Ctx) {
	// 	msg, _ = services.GetAllPerson()
	// 	c.Status(200).JSON(msg)
	// })
	app.Get("/test", func(c *fiber.Ctx) error {
		msg, err := services.GetAllPerson()
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		return c.Status(200).JSON(msg)
	})
	// 	person_entry_point := app.Group("/person")
	// 	PersonRoute(person_entry_point)
}
