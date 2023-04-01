package main

import (
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// New fiber app
	app := fiber.New()

	// http middleware -> fiber.Handler
	app.Use(adaptor.HTTPMiddleware(logMiddleware))

	// Listen on port 3000
	app.Listen(":3000")
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("log middleware")
		next.ServeHTTP(w, r)
	})
}
