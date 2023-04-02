package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	// app.Use(cache.New())

	// recover from panic
	app.Use(recover.New())

	app.Use(logger.New())
	// if os.Getenv("ENV") == "production" {
	// 	// app.Use(helmet.New())
	// 	// app.Use(limiter.New())
	// 	// app.Use(requestid.New())
	// 	// app.Use(timeout.New())
	// 	// app.Use(proxy.New())
	// }

	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// app.Get("/ocr", func(c *fiber.Ctx) error {

	// 	imagePath := c.Query("image")
	// 	if imagePath == "" {
	// 		return c.SendString("Image path is required")
	// 	}

	// 	client := gosseract.NewClient()
	// 	defer client.Close()
	// 	client.SetImage("test.png")
	// 	text, _ := client.Text()
	// 	return c.SendString(text)
	// })

	app.Get("/test2", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
