package main

import (
	"net/http"

	"github.com/gofiber/adaptor"
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

	app.Listen(3000)
}

// handler is the entry point for the Cloud Function.
func handler(w http.ResponseWriter, r *http.Request) {
	adaptor.HTTPHandler(app)(w, r)
}
