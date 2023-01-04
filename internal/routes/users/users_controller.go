package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/middleware"
	"github.com/kyh0703/stock-server/internal/routes/users/dtos"
	"github.com/kyh0703/stock-server/internal/types"
)

type UsersController struct {
	usersService *UsersService
}

func NewUsersController(usersService *UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (ctrl *UsersController) Path() string {
	return "users"
}

func (ctrl *UsersController) Index(router fiber.Router) {
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
func (ctrl *UsersController) Register(c *fiber.Ctx) error {
	req := new(dtos.UsersRegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctrl.usersService.Register(c, req)
}

// Login        godoc
// @Summary     login jwt users
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/login [post]
func (ctrl *UsersController) Login(c *fiber.Ctx) error {
	req := new(dtos.UsersLoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersService.Login(c, req)
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/profile [get]
func (ctrl *UsersController) Profile(c *fiber.Ctx) error {
	req := new(dtos.UsersProfileRequest)
	req.ID = c.Context().UserValue(types.ContextKeyUserID).(int)

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersService.GetUserDetail(c, req)
}

// Logout       godoc
// @Summary     Show a account
// @Description get string by ID
// @ID          get-string-by-int
// @Accept      json
// @Produce     json
// @Router      /users/logout [post]
func (ctrl *UsersController) Logout(c *fiber.Ctx) error {
	token := c.Context().UserValue(types.ContextKeyAccessToken).(string)
	return ctrl.usersService.Logout(c, token)
}

// Refresh      godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/refresh [post]
func (ctrl *UsersController) Refresh(c *fiber.Ctx) error {
	req := new(dtos.UsersRefreshRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, fiber.ErrBadRequest)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersService.RefreshToken(c, req)
}
