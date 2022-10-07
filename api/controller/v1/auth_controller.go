package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	path string
	rg   *gin.RouterGroup
}

func NewAuthController(rg *gin.RouterGroup) *authController {
	return &authController{
		path: "/auth",
		rg:   rg,
	}
}

func (ctrl *authController) Route() *gin.RouterGroup {
	route := ctrl.rg.Group(ctrl.path)
	route.POST("/register", ctrl.register)
	route.POST("/login", ctrl.login)
	route.GET("/check", ctrl.check)
	route.POST("/logout", ctrl.logout)
	return route
}

func (ctrl *authController) register(c *gin.Context) {
	// define body in request message
	type request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=3,max=10"`
	}
	// validator
	req := new(request)
	if err := c.Bind(req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// process request
	c.Status(http.StatusOK)
}

func (ctrl *authController) login(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (ctrl *authController) check(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (ctrl *authController) logout(c *gin.Context) {
	c.Status(http.StatusOK)
}
