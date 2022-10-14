package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/middleware"
	v1 "github.com/kyh0703/stock-server/routes/v1"
)

func New() *fiber.App {
	// create app
	app := fiber.New(config.FiberConfig())

	// middleware
	app.Use(recover.New())
	app.Use(middleware.SetJSON())

	// swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// routes
	api := app.Group("/api")
	v1r := api.Group("v1")
	{
		v1.NewAuthController(v1r).Index()
		v1.NewPostController(v1r).Index()
	}
	return app
}
