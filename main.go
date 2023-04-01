package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	// Connect to the Database
	// database.ConnectDB()

	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	// app.Use(cache.New())

	// recover from panic
	app.Use(recover.New())

	app.Use(logger.New())

	if os.Getenv("ENV") == "production" {
		// app.Use(helmet.New())
		// app.Use(limiter.New())
		// app.Use(requestid.New())
		// app.Use(timeout.New())
		// app.Use(proxy.New())
	}

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("localhost:" + port))
}
