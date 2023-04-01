package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

var (
	app *fiber.App
)

func init() {
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// app.Listen(8080)
}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.Listen(3000)
}
