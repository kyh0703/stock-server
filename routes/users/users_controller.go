package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/auth"
	usersdto "github.com/kyh0703/stock-server/routes/users/dto"
	"github.com/kyh0703/stock-server/types"
)

type usersController struct {
	path        string
	userService UsersService
	authService auth.AuthService
}

func NewUsersController() *usersController {
	return &usersController{
		path: "users",
	}
}

func (ctrl *usersController) Path() string {
	return ctrl.path
}

func (ctrl *usersController) Routes(router fiber.Router) {
	router.Post("/register", ctrl.Register)
	router.Post("/login", ctrl.Login)
	router.Get("/profile", middleware.TokenAuth(), ctrl.Profile)
	router.Post("/logout", middleware.TokenAuth(), ctrl.Logout)
	router.Post("/refresh", ctrl.Refresh)
}

// SignUp       godoc
// @Summary     SignUp auth info
// @Description SignUp stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/register [post]
func (ctrl *usersController) Register(c *fiber.Ctx) error {
	var (
		req usersdto.UserRegisterRequest
		res usersdto.UserRegisterResponse
	)

	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// validate
	if err := validator.New().StructCtx(c.Context(), res); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// compare "Password" to "PasswordConfirm"
	if req.Password != req.PasswordConfirm {
		return c.App().ErrorHandler(c, types.ErrPasswordNotCompareConfirm)
	}

	// hash password
	hash, err := ctrl.authService.HashPassword(req.Password)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	// check the exist user
	exist, err := ctrl.userService.IsExistEmail(c.Context(), req.Email)
	if err != nil || exist {
		if exist {
			return c.App().ErrorHandler(c, types.ErrUserExist)
		} else {
			return c.App().ErrorHandler(c, types.ErrServerInternal)
		}
	}

	// register in database
	if _, err := ctrl.userService.SaveUser(c.Context(), req.Username, req.Email, hash); err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	// response token data
	return c.JSON(res)
}

// Login        godoc
// @Summary     login jwt users
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/login [post]
func (ctrl *usersController) Login(c *fiber.Ctx) error {
	var (
		req usersdto.UserLoginRequest
		res usersdto.UserLoginResponse
	)

	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// validate request message
	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// check exist user from database
	user, err := ctrl.userService.FindByEmail(c.Context(), req.Email)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserNotExist)
	}

	// verify password
	ok, err := ctrl.authService.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		return c.App().ErrorHandler(c, types.ErrPasswordInvalid)
	}

	// save jwt auth
	token, err := ctrl.authService.Login(user.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	// set refresh token in cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token.RefreshToken
	cookie.HTTPOnly = true
	cookie.Secure = true
	c.Cookie(cookie)

	// response token data
	res.ID = user.ID
	res.Email = user.Email
	res.Username = user.Username
	res.AccessToken = token.AccessToken
	res.AccessTokenExpires = token.AccessTokenExpires
	return c.JSON(res)
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/profile [get]
func (ctrl *usersController) Profile(c *fiber.Ctx) error {
	var (
		userId int
		res    usersdto.UserProfileResponse
	)
	userId = c.UserContext().Value(ContextKeyUserID).(int)
	user, err := ctrl.userService.FindOne(c.Context(), userId)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserNotExist)
	}
	res.ID = user.ID
	res.Email = user.Email
	res.Username = user.Username
	return c.JSON(res)
}

// Logout       godoc
// @Summary     Show a account
// @Description get string by ID
// @ID          get-string-by-int
// @Accept      json
// @Produce     json
// @Router      /users/logout [post]
func (ctrl *usersController) Logout(c *fiber.Ctx) error {
	token := c.UserContext().Value(ContextKeyAccessToken).(string)
	if err := ctrl.authService.Logout(token); err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}
	return c.SendStatus(http.StatusNoContent)
}

// Refresh      godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/refresh [post]
func (ctrl *usersController) Refresh(c *fiber.Ctx) error {
	var (
		req usersdto.RefreshTokenRequest
		res usersdto.RefreshTokenResponse
	)

	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.App().ErrorHandler(c, fiber.ErrBadRequest)
	}

	// validate request message
	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// refresh token
	token, err := ctrl.authService.Refresh(req.RefreshToken)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	res.AccessToken = token.AccessToken
	return c.JSON(res)
}
