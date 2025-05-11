package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/srohatgi01/tpp-helper/data"
	"github.com/srohatgi01/tpp-helper/handlers"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./public/views", ".html")
	engine.AddFunc("title", func(s string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(s[:1]) + s[1:]
	})
	engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from ./public
	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Papers": data.GetPaperList(),
		})
	})

	app.Post("/estimate", handlers.EstimateHandler)

	app.Listen(":3000")
}
