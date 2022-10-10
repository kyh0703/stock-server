package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/ent"
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

func (ctrl *authController) Index() *gin.RouterGroup {
	route := ctrl.rg.Group(ctrl.path)
	route.POST("/register", ctrl.Register)
	route.POST("/login", ctrl.Login)
	route.GET("/check", ctrl.Check)
	route.POST("/logout", ctrl.Logout)
	return route
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Param       username string
// @Param       password string
// @Success     200
// @Router      /auth/register [post]
func (ctrl *authController) Register(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// validator
	req := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=3,max=10"`
	}{}
	if err := c.Bind(req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// register
	user, err := db.User.Create().
		SetUsername(req.Username).
		Save(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user)
}

// Login        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/login [post]
func (ctrl *authController) Login(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/check [get]
func (ctrl *authController) Check(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Logout       godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/logout [post]
func (ctrl *authController) Logout(c *gin.Context) {
	c.Status(http.StatusOK)
}
