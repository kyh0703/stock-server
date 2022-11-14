package posts

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/posts/dto"
	"github.com/kyh0703/stock-server/types"
)

type postController struct {
	path     string
	postsSvc postsService
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
	router.Delete("/:id", middleware.TokenAuth(), ctrl.RemovePostById)
}

// Write        godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/write [post]
func (ctrl *postController) Write(c *fiber.Ctx) error {
	req := new(dto.PostCreateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	req.UserID = c.UserContext().Value("user_id").(int)
	return ctrl.postsSvc.SavePost(c, req)
}

// List         godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts [get]
func (ctrl *postController) List(c *fiber.Ctx) error {
	var (
		page     = c.Query("page", "1")
		limit    = c.Query("limit", "10")
		tag      = c.Query("tag")
		username = c.Query("username")
	)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	req := dto.PostListRequest{
		Page:     pageInt,
		Limit:    limitInt,
		Tag:      tag,
		Username: username,
	}
	return ctrl.postsSvc.GetPosts(c, req)
}

// GetPostById  godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/:id [post]
func (ctrl *postController) GetPostById(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	req := dto.PostFetchRequest{
		ID: postId,
	}

	return ctrl.postsSvc.GetPost(c, req)
}

// GetPostById  godoc
// @Summary     remove post
// @Description remove post api
// @Tags        post
// @Produce     json
// @Success     200
// @Router      /posts/:id [delete]
func (ctrl *postController) RemovePostById(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	return ctrl.postsSvc.RemovePost(c, postId)
}
