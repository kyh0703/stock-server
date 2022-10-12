package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/api/auth"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/lib"
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
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=3,max=10"`
		Username string `json:"username" binding:"required"`
	}{}
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// hash password
	hashPassword, err := lib.HashPassword(req.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// check the exist user
	exist, err := db.User.
		Query().
		Where(
			user.EmailContains(req.Email),
		).Exist(c)
	if err != nil || exist {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	// register in database
	user, err := db.User.
		Create().
		SetUsername(req.Username).
		SetPassword(hashPassword).
		SetEmail(req.Email).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// access token
	accessToken, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.SetCookie("access-token", accessToken, 60*60*24*7, "/", "localhost:8000", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"message":     "회원가입 완료",
		"accessToken": accessToken,
	})
}

// Login        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Param       username string
// @Param       password string
// @Success     200
// @Router      /auth/login [post]
func (ctrl *authController) Login(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// validator
	req := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=3,max=10"`
	}{}
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// check exist user
	user, err := db.User.
		Query().
		Where(
			user.EmailContains(req.Email)).
		Only(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// verify password
	ok, err := lib.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// access token
	accessToken, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// set cookie
	c.SetCookie("access-token", accessToken, 60*60*24*7, "/", "localhost:8000", false, true)
	// TODO (refreshed token) 추후 작업 예정입니다.
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"message":     "토큰 발급 완료",
		"accessToken": accessToken,
	})
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/check [get]
func (ctrl *authController) Check(c *gin.Context) {
	userID := c.Request.Header.Get("x-request-id")
	if userID == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
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
	c.SetCookie("access-token", "", 0, "/", "", false, true)
	c.Status(http.StatusNoContent)
}
