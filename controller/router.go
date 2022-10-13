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

func NewRouter(ec *ent.Client, rc *redis.Client) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.SetDatabase(ec))
	router.Use(middleware.SetRedis(rc))
	router.Use(middleware.SetJSON())
	return router
}

func SetupRouter(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	rg := api.Group("v1")
	{
		v1.NewAuthController(rg).Index()
		v1.NewPostController(rg).Index()
	}
}
