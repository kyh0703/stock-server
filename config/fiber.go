package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func FiberConfig() fiber.Config {
	// Initialize standard Go Html template engine
	engine := html.New("./views", ".html")

	// Debug will print each template that is parsed
	engine.Debug(true)

	// Reload the templates on each render, good for development
	engine.Reload(true)

	// Define server settings.
	return fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		ReadTimeout: Env.ReadTimeout,
	}
}
