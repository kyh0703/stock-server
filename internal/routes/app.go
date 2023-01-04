package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/kyh0703/stock-server/configs"
	"github.com/kyh0703/stock-server/internal/middleware"
	"github.com/kyh0703/stock-server/internal/types"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/kyh0703/stock-server/internal/routes/posts"
	_ "github.com/kyh0703/stock-server/internal/routes/users"
)

var AppModule types.Module

func init() {
}

func SetupApp() *fiber.App {
	// create app
	app := fiber.New(configs.Fiber(false))

	// middleware
	app.Use(middleware.SetUserContext())
	app.Use(cors.New())
	// app.Use(recover.New())
	app.Use(logger.New(configs.Logger()))

	// monitor
	app.Get("/metrics", monitor.New(configs.Monitor()))

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// set fiber engine
	AppModule.Init(app)
	return app
}
