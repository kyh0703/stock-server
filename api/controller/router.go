package controller

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/kyh0703/stock-server/api/controller/v1"
	"github.com/kyh0703/stock-server/api/middleware"
	"github.com/kyh0703/stock-server/ent"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.SetDatabase(client))
	router.Use(middleware.SetJSON())
	router.Use(middleware.SetAuthentication())
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
