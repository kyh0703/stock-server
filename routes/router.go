package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/kyh0703/stock-server/config"
	v1 "github.com/kyh0703/stock-server/routes/v1"
)

func New() *fiber.App {
	// create app
	app := fiber.New(config.Fiber(false))

	// middleware
	app.Use(cors.New())
	// app.Use(recover.New())
	app.Use(logger.New(config.Logger()))

	// monitor
	app.Get("/metrics", monitor.New(config.Monitor()))

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// routes
	api := app.Group("/api")
	v1api := api.Group("v1")
	{
		v1.NewAuthController(v1api).Index()
		v1.NewPostController(v1api).Index()
	}
	return app
}
