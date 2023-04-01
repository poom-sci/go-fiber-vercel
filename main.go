package main

import (
	"github.com/gofiber/adaptor"
	"github.com/gofiber/fiber"
)

// var (
// 	app *fiber.App
// )

// func init() {
// 	app = fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send("Hello, World!")
// 	})

// 	// app.Listen(3000)
// }

// ADD THIS SCRIPT
func Handler(c *fiber.Ctx) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	adaptor.HTTPHandler(app).ServeHTTP(c)

}
