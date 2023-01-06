package posts

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/app/posts/dtos"
	"github.com/kyh0703/stock-server/internal/middleware"
	"github.com/kyh0703/stock-server/internal/types"

	postsdto "github.com/kyh0703/stock-server/internal/app/posts/dtos"
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
	router.Get("/", ctrl.FindPagePost)
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

	if res, err := ctrl.postsService.Create(c, req); err != nil {
		return err
	} else {
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

// List         godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts [get]
func (ctrl *PostsController) FindPagePost(c *fiber.Ctx) error {
	req := new(dtos.PagePostsDto)
	if err := c.QueryParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	posts, err := ctrl.postsService.PagePosts(c, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(posts)
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

	post, err := ctrl.postsService.FindOne(c, id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(post)
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

	post, err := ctrl.postsService.Update(c, id, dto)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(post)
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

	if err := ctrl.postsService.RemovePost(c, id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
