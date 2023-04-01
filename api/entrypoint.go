package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
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
	// Create a new fasthttp.RequestCtx object
	req := new(fasthttp.RequestCtx)
	req.Request.SetRequestURI(r.RequestURI)
	req.Request.Header.SetMethod(r.Method)
	// r.Header.VisitAll(func(key, value []byte) {
	// 	req.Request.Header.Set(key, value)
	// })
	for k, v := range r.Header {
		req.Request.Header.Set(k, v[0])
	}
	// req.Request.SetBodyStream(r.Body, r.ContentLength)
	// req.Request.SetBody(r.Body)

	app.Handler()(req)

	for k, v := range req.Response.Header.Header() {
		// w.Header().Set(k, string(v))
		w.Header().Set(string(rune(k)), string(v))
	}
	w.WriteHeader(req.Response.StatusCode())
	w.Write(req.Response.Body())

}

func main() {
	err := http.ListenAndServe(":3000", http.HandlerFunc(Handler))
	if err != nil {
		panic(err)
	}
}
