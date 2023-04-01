package main

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func main() {
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {

	http.ListenAndServe(":3000", adaptor.FiberApp(app))

}
