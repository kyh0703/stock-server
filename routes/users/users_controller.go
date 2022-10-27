package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/auth"
	usersdto "github.com/kyh0703/stock-server/routes/users/dto"
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
	router.Post("/signup", ctrl.SignUp)
	router.Post("/login", ctrl.Login)
	router.Get("/check", middleware.TokenAuth(), ctrl.Check)
	router.Post("/logout", middleware.TokenAuth(), ctrl.Logout)
	router.Post("/refresh", ctrl.Refresh)
}

// SignUp       godoc
// @Summary     SignUp auth info
// @Description SignUp stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/signup [post]
func (ctrl *usersController) SignUp(c *fiber.Ctx) error {
	// body parser
	var dto usersdto.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// compare "Password" to "PasswordConfirm"
	if dto.Password != dto.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).SendString("password not equal confirm")
	}
	// hash password
	hash, err := ctrl.authService.HashPassword(dto.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// check the exist user
	exist, err := ctrl.userService.IsExistEmail(c.Context(), dto.Email)
	if err != nil || exist {
		if exist {
			return fiber.NewError(fiber.StatusConflict, "사용자가 이미 존재합니다")
		} else {
			return fiber.NewError(fiber.StatusConflict, "잠시후 다시 이용하여 주시기 바랍니다")l
		}
	}
	// register in database
	if _, err := ctrl.userService.SaveUser(c.Context(), dto.Username, dto.Email, hash); err != nil {
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
// @Router      /users/login [post]
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
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// set cookie data
	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token.AccessToken
	cookie.HTTPOnly = true
	c.Cookie(cookie)
	// response token data
	return c.JSON(fiber.Map{
		"user":         user.Email,
		"access_token": token.AccessToken,
	})
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/check [get]
func (ctrl *usersController) Check(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// Logout       godoc
// @Summary     Show a account
// @Description get string by ID
// @ID          get-string-by-int
// @Accept      json
// @Produce     json
// @Router      /users/logout [post]
func (ctrl *usersController) Logout(c *fiber.Ctx) error {
	token := c.UserContext().Value("token").(string)
	if err := ctrl.authService.Logout(token); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
	var dto usersdto.RefreshTokenDTO
	// body parser
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// validate request message
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	token, err := ctrl.authService.Refresh(dto.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(token)
}
