package posts

import (
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	postsdto "github.com/kyh0703/stock-server/routes/posts/dto"
	"github.com/kyh0703/stock-server/types"
)

type postController struct {
	path         string
	postsService postsService
}

func NewPostController() *postController {
	return &postController{
		path: "posts",
	}
}

func (ctrl *postController) Path() string {
	return ctrl.path
}

func (ctrl *postController) Routes(router fiber.Router) {
	router.Get("/", ctrl.List)
	router.Get("/:id", ctrl.GetPostById)
	router.Post("/write", middleware.TokenAuth(), ctrl.Write)
}

// Write        godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/write [post]
func (ctrl *postController) Write(c *fiber.Ctx) error {
	// body parser
	var dto postsdto.CreatePostDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// validate request message
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// set user id
	dto.UserID = c.UserContext().Value("user_id").(int)

	// save the database
	post, err := ctrl.postsService.SavePost(c.Context(), dto)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

// List         godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts [get]
func (ctrl *postController) List(c *fiber.Ctx) error {
	// get query data
	var (
		page = c.Query("page", "1")
		tag  = c.Query("tag")
		name = c.Query("username")
	)

	// parse page
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// select data
	posts, err := ctrl.postsService.FindPagesByNameOrTag(c.Context(), tag, name, pageInt, 10)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	// get post count
	count, err := ctrl.postsService.GetCountByNameOrTag(c.Context(), tag, name)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	c.Response().Header.Set("last-page", strconv.Itoa(int(math.Ceil(float64(count/10)))))
	return c.Status(fiber.StatusOK).JSON(posts)
}

// GetPostById  godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/:id [post]
func (ctrl *postController) GetPostById(c *fiber.Ctx) error {
	// validate
	postId, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	// get post data
	post, err := ctrl.postsService.FindOne(c.Context(), postId)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
