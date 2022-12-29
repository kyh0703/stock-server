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

	// set error handler
	cfg.ErrorHandler = func(c *fiber.Ctx, err error) error {
		// Status code defaults to 500
		var (
			code    = fiber.ErrInternalServerError.Code
			message = fiber.ErrInternalServerError.Message
		)

		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		// set json data
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		return c.Status(code).JSON(fiber.Map{
			"statusCode": code,
			"message":    message,
		})
	}

	return cfg
}
