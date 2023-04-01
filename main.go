package main

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

func main() {
	app := fiber.New()

	// app.Get("/greet", greet)

	// Listen on port 3000
	http.ListenAndServe(":3000", adaptor.FiberApp(app))
}

func greet(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// add fiber app to handle request
	// http.ListenAndServe(":3000", adaptor.FiberApp(app))
	app.Get("/greet", greet)

}
