package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/routes/posts"
	"github.com/kyh0703/stock-server/routes/users"
	"github.com/kyh0703/stock-server/types"
)

var module types.Module

func SetUpRouter() *fiber.App {
	// create app
	app := fiber.New(config.Fiber(false))

	// middleware
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(config.Logger()))

	// monitor
	app.Get("/metrics", monitor.New(config.Monitor()))

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// set fiber engine
	module.SetEngine(app)

	// controller
	module.AttachController(users.NewUsersController(app))
	module.AttachController(posts.NewPostController(app))
	return app
}
