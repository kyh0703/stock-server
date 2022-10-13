package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/lib/jwt"
	"github.com/kyh0703/stock-server/lib/password"
	"github.com/kyh0703/stock-server/middleware"
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
	route.GET("/check", middleware.TokenAuth(), ctrl.Check)
	route.POST("/logout", middleware.TokenAuth(), ctrl.Logout)
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
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required,min=3,max=10"`
		PasswordConfirm string `json:"passwordConfirm" binding:"required,min=3,max=10"`
		Username        string `json:"username" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// compare "Password" to "PasswordConfirm"
	if req.Password != req.PasswordConfirm {
		c.AbortWithError(http.StatusBadRequest, errors.New("confirm password"))
		return
	}
	// hash password
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// check the exist user
	exist, err := db.User.
		Query().
		Where(user.EmailContains(req.Email)).
		Exist(c)
	if err != nil || exist {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	// register in database
	_, err = db.User.
		Create().
		SetUsername(req.Username).
		SetPassword(hashPassword).
		SetEmail(req.Email).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
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
	rc, _ := c.Keys["redis"].(*redis.Client)
	// validator
	req := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=3,max=10"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// check exist user from database
	user, err := db.User.
		Query().
		Where(user.EmailContains(req.Email)).
		Only(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// verify password
	ok, err := password.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// create token
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	// save the redis
	if err := jwt.SaveTokenData(rc, user.ID, token); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	// response token data
	c.JSON(http.StatusOK, map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
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
	rc, _ := c.Keys["redis"].(*redis.Client)
	// delete token data
	au, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	deleted, err := jwt.DeleteTokenData(rc, au.AccessUUID)
	if err != nil || deleted == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.SetCookie("access-token", "", 0, "/", "", false, true)
	c.Status(http.StatusNoContent)
}
