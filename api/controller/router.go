package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	v1 "github.com/kyh0703/stock-server/api/controller/v1"
)

type Controller interface {
	Route()
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	return router
}

func Index(router *gin.Engine) {
	rg := router.Group("/api/v1")
	v1.NewAuthController(rg).Route()
	log.Println("hgihihihi")
}
