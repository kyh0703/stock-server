package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Fiber(existView bool) fiber.Config {
	cfg := fiber.Config{}
	// Initialize standard Go Html template engine
	if existView {
		engine := html.New("./views", ".html")
		// Debug will print each template that is parsed
		engine.Debug(true)
		// Reload the templates on each render, good for development
		engine.Reload(true)
		// Set cfg
		cfg.Views = engine
		cfg.ViewsLayout = "layouts/main"
	}
	// Define server settings.
	return cfg
}
