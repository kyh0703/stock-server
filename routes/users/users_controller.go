package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/users/dto"
	"github.com/kyh0703/stock-server/types"
)

type usersController struct {
	path     string
	usersSvc UsersService
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
	req := new(dto.UsersRegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctrl.usersSvc.Register(c, req)
}

// Login        godoc
// @Summary     login jwt users
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/login [post]
func (ctrl *usersController) Login(c *fiber.Ctx) error {
	req := new(dto.UsersLoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersSvc.Login(c, req)
}

// Check        godoc
// @Summary     Get books array
// @Description Responds with the list of all books as JSON.
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/profile [get]
func (ctrl *usersController) Profile(c *fiber.Ctx) error {
	req := new(dto.UsersProfileRequest)
	req.ID = c.UserContext().Value("user_id").(int)

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersSvc.GetUserDetail(c, req)
}

// Logout       godoc
// @Summary     Show a account
// @Description get string by ID
// @ID          get-string-by-int
// @Accept      json
// @Produce     json
// @Router      /users/logout [post]
func (ctrl *usersController) Logout(c *fiber.Ctx) error {
	token := c.UserContext().Value("access_token").(string)
	return ctrl.usersSvc.Logout(c, token)
}

// Refresh      godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /users/refresh [post]
func (ctrl *usersController) Refresh(c *fiber.Ctx) error {
	req := new(dto.UsersRefreshRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, fiber.ErrBadRequest)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.usersSvc.RefreshToken(c, req)
}
