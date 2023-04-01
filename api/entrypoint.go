package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

// CREATE ENDPOINT
func myRoute(r fiber.Router) {
	r.Get("/admin", func(c *fiber.Ctx) error {
		return c.SendString("Hello from golang in vercel")
	})
}

func init() {
	app = fiber.New()
	r := app.Group("/")
	myRoute(r)
}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.Handler()(w, r)
}

func main() {
	err := http.ListenAndServe(":3000", http.HandlerFunc(Handler))
	if err != nil {
		panic(err)
	}
}
