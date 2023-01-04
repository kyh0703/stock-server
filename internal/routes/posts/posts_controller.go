package posts

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/middleware"
	"github.com/kyh0703/stock-server/internal/routes/posts/dtos"
	postsdto "github.com/kyh0703/stock-server/internal/routes/posts/dtos"
	"github.com/kyh0703/stock-server/internal/types"
)

type PostsController struct {
	postsService *PostsService
}

func NewPostController(postsService *PostsService) *PostsController {
	return &PostsController{
		postsService: postsService,
	}
}

func (ctrl *PostsController) Path() string {
	return "posts"
}

func (ctrl *PostsController) Index(router fiber.Router) {
	router.Post("/write", middleware.TokenAuth(), ctrl.CreatePost)
	router.Get("/", ctrl.FindAllPost)
	router.Get("/:id", ctrl.FindPost)
	router.Patch("/:id", middleware.TokenAuth(), ctrl.UpdatePost)
	router.Delete("/:id", middleware.TokenAuth(), ctrl.RemovePost)
}

// Write        godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/write [post]
func (ctrl *PostsController) CreatePost(c *fiber.Ctx) error {
	req := new(postsdto.CreatePostDto)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	req.UserID = c.Context().UserValue(types.ContextKeyUserID).(int)
	return ctrl.postsService.Create(c, req)
}

// List         godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts [get]
func (ctrl *PostsController) FindAllPost(c *fiber.Ctx) error {
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

	req := &dtos.FindPostsDto{
		Page:     pageInt,
		Limit:    limitInt,
		Tag:      tag,
		Username: username,
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsService.GetPosts(c, req)
}

// GetPostById  godoc
// @Summary     fetch posts api
// @Description fetch posts by postID
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/:id [post]
func (ctrl *PostsController) FindPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsService.FindOne(c, id)
}

// UpdatePostById godoc
// @Summary       update post api
// @Description   update post by id
// @Tags          post
// @Produce       json
// @Success       200
// @Router        /posts [patch]
func (ctrl *PostsController) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	dto := new(dtos.CreatePostDto)
	if err := c.BodyParser(dto); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), dto); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsService.Update(c, id, dto)
}

// RemovePostById godoc
// @Summary       remove post api
// @Description   remove post by id
// @Tags          post
// @Produce       json
// @Success       200
// @Router        /posts [delete]
func (ctrl *PostsController) RemovePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsService.RemovePost(c, id)
}
