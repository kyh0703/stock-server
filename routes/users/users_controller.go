package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/lib/jwt"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/auth"
	usersdto "github.com/kyh0703/stock-server/routes/users/dto"
)

type usersController struct {
	path        string
	router      fiber.Router
	userService UsersService
	authService auth.AuthService
}

func NewUsersController(router fiber.Router) *usersController {
	return &usersController{
		path:   "auth",
		router: router,
	}
}

func (ctrl *usersController) Path() string {
	return ctrl.path
}

func (ctrl *usersController) Routes() {
	ctrl.router.Post("/signup", ctrl.SignUp)
	ctrl.router.Post("/login", ctrl.Login)
	ctrl.router.Get("/check", middleware.TokenAuth(), ctrl.Check)
	ctrl.router.Post("/logout", middleware.TokenAuth(), ctrl.Logout)
	ctrl.router.Post("/refresh", ctrl.Refresh)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /user/signup [post]
func (ctrl *usersController) SignUp(c *fiber.Ctx) error {
	var dto usersdto.CreateUserDTO
	// body parser
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// compare "Password" to "PasswordConfirm"
	if dto.Password != dto.PasswordConfirm {
		return fiber.NewError(fiber.StatusBadRequest, "password not equal passwordConfirm")
	}
	// hash password
	hash, err := ctrl.authService.HashPassword(dto.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// check the exist user
	if _, err := ctrl.userService.FindByEmail(c.Context(), dto.Email); err != nil {
		return fiber.ErrConflict
	}
	// register in database
	if _, err := ctrl.userService.SaveUser(c.Context(), dto.Name, dto.Email, hash); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}

// Login        godoc
// @Summary     login jwt users
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/login [post]
func (ctrl *usersController) Login(c *fiber.Ctx) error {
	var dto usersdto.UserLoginDTO
	// body parser
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate request message
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	// check exist user from database
	user, err := ctrl.userService.FindByEmail(c.Context(), dto.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// verify password
	ok, err := ctrl.authService.CompareHashPassword(user.Password, dto.Password)
	if err != nil || !ok {
		return fiber.ErrUnauthorized
	}
	// save jwt auth
	token, err := ctrl.authService.Login(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// response token data
	return c.JSON(token)
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/check [get]
func (ctrl *usersController) Check(c *fiber.Ctx) error {
	userID := c.UserContext().Value("user_id").(int)
	if userID == 0 {
		return fiber.ErrBadRequest
	}
	return c.SendStatus(fiber.StatusOK)
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
func (ctrl *usersController) Logout(c *fiber.Ctx) error {
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
func (ctrl *usersController) Refresh(c *fiber.Ctx) error {
	req := struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}{}
	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate request message
	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
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
