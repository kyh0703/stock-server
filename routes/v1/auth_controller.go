package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/lib/jwt"
	"github.com/kyh0703/stock-server/lib/password"
	"github.com/kyh0703/stock-server/middleware"
)

type authController struct {
	path   string
	router fiber.Router
}

func NewAuthController(router fiber.Router) *authController {
	return &authController{
		path:   "/auth",
		router: router,
	}
}

func (ctrl *authController) Index() fiber.Router {
	r := ctrl.router.Group(ctrl.path)
	r.Post("/register", ctrl.Register)
	r.Post("/login", ctrl.Login)
	r.Get("/check", middleware.TokenAuth(), ctrl.Check)
	r.Post("/logout", middleware.TokenAuth(), ctrl.Logout)
	r.Post("/refresh", ctrl.Refresh)
	return r
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/register [post]
func (ctrl *authController) Register(c *fiber.Ctx) error {
	req := struct {
		Email           string `json:"email" validate:"required"`
		Password        string `json:"password" validate:"required,min=3,max=10"`
		PasswordConfirm string `json:"passwordConfirm" validate:"required,min=3,max=10"`
		Username        string `json:"username" validate:"required"`
	}{}
	// body parser
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate
	if errors := validator.New().StructCtx(c.Context(), req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	// compare "Password" to "PasswordConfirm"
	if req.Password != req.PasswordConfirm {
		return fiber.ErrBadRequest
	}
	// hash password
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// check the exist user
	exist, err := database.Ent().User.
		Query().
		Where(user.EmailContains(req.Email)).
		Exist(c.Context())
	if err != nil || exist {
		return fiber.ErrConflict
	}
	// register in database
	user, err := database.Ent().User.
		Create().
		SetUsername(req.Username).
		SetPassword(hashPassword).
		SetEmail(req.Email).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// create token
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// save the redis
	if err := jwt.SaveTokenData(user.ID, token); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}

// Login        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/login [post]
func (ctrl *authController) Login(c *fiber.Ctx) error {
	req := struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required,min=3,max=10"`
	}{}
	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate request message
	if errors := validator.New().StructCtx(c.Context(), req); errors != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errors)
	}
	// check exist user from database
	user, err := database.Ent().User.
		Query().
		Where(user.EmailContains(req.Email)).
		Only(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// verify password
	ok, err := password.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		return fiber.ErrUnauthorized
	}
	// create token
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// save the redis
	if err := jwt.SaveTokenData(user.ID, token); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// response token data
	return c.JSON(fiber.Map{
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
func (ctrl *authController) Check(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
func (ctrl *authController) Logout(c *fiber.Ctx) error {
	// get token metadata
	accessUser, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	// delete token data
	deleted, err := jwt.DeleteTokenData(accessUser.AccessUUID)
	if err != nil || deleted == 0 {
		return fiber.ErrUnauthorized
	}
	return c.SendStatus(http.StatusNoContent)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/refresh [post]
func (ctrl *authController) Refresh(c *fiber.Ctx) error {
	// validate request message
	req := struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// Get claims with refresh secret key
	claims, err := jwt.ValidateToken(c, req.RefreshToken, config.Env.RefreshSecretKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// token
	refreshUUID, ok := claims["refresh_uuid"].(string) // convert the interface to string
	if !ok {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	deleted, err := jwt.DeleteTokenData(refreshUUID)
	if err != nil || deleted == 0 {
		return fiber.ErrUnprocessableEntity
	}
	token, err := jwt.CreateToken(int(userID))
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	if err := jwt.SaveTokenData(int(userID), token); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}
