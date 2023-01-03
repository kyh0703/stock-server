package routes

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/kyh0703/stock-server/configs"
	"github.com/kyh0703/stock-server/internal/middleware"
	"github.com/kyh0703/stock-server/internal/types"
)

var AppModule types.Module

func init() {
	fmt.Println("hihihihihih print App")
}

func NewApp() *fiber.App {
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
	AppModule.Engine = app

	// controller
	AppModule.Init()
	return app
}
