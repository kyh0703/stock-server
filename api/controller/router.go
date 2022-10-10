package controller

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/kyh0703/stock-server/api/controller/v1"
	"github.com/kyh0703/stock-server/api/middleware"
	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/ent"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	Route()
}

func NewRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middleware.SetDatabase(client))
	switch config.Env.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}
	return router
}

func SetupRouter(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rg := router.Group("/api/v1")
	v1.NewAuthController(rg).Index()
}
