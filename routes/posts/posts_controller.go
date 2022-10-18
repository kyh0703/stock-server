package posts

import (
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	postsdto "github.com/kyh0703/stock-server/routes/posts/dto"
)

type postController struct {
	path         string
	postsService postsService
}

func NewPostController() *postController {
	return &postController{
		path: "post",
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

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) Write(c *fiber.Ctx) error {
	var dto postsdto.CreatePostDTO
	// body parser
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// validate request message
	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	// set user id
	userId := c.UserContext().Value("user_id").(int)
	if userId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "user id dose not exist")
	}
	dto.UserID = userId
	// save the database
	post, err := ctrl.postsService.SavePost(c.Context(), dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/register [post]
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
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// select data
	posts, err := ctrl.postsService.FindPagesByNameOrTag(c.Context(), tag, name, pageInt, 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// get post count
	count, err := ctrl.postsService.GetCountByNameOrTag(c.Context(), tag, name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	c.Response().Header.Set("last-page", strconv.Itoa(int(math.Ceil(float64(count/10)))))
	return c.Status(fiber.StatusOK).JSON(posts)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) GetPostById(c *fiber.Ctx) error {
	// validate
	postId, err := c.ParamsInt("id", 0)
	if err != nil {
		return fiber.ErrBadRequest
	}
	// get post data
	post, err := ctrl.postsService.FindOne(c.Context(), postId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}
