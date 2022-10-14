package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	v1 "github.com/kyh0703/stock-server/controller/v1"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(ec *ent.Client, rc *redis.Client) *gin.Engine {
	router := gin.Default()
	// set middleware
	router.Use(middleware.SetEntClient(ec))
	router.Use(middleware.SetRedisClient(rc))
	router.Use(middleware.SetJSON())
	// set swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// routes
	api := router.Group("/api")
	rg := api.Group("v1")
	{
		v1.NewAuthController(rg).Index()
		v1.NewPostController(rg).Index()
	}
	return router
}
