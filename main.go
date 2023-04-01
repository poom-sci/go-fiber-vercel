package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}
