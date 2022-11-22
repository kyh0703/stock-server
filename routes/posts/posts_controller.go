package posts

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/middleware"
	"github.com/kyh0703/stock-server/routes/posts/dto"
	"github.com/kyh0703/stock-server/types"
)

type postsController struct {
	path     string
	postsSvc postsService
}

func NewPostController() *postsController {
	return &postsController{
		path: "posts",
	}
}

func (ctrl *postsController) Path() string {
	return ctrl.path
}

func (ctrl *postsController) Routes(router fiber.Router) {
	router.Post("/write", middleware.TokenAuth(), ctrl.Write)
	router.Get("/", ctrl.List)
	router.Get("/:id", ctrl.GetPostById)
	router.Patch("/", middleware.TokenAuth(), ctrl.UpdatePostById)
	router.Delete("/", middleware.TokenAuth(), ctrl.RemovePostById)
}

// Write        godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/write [post]
func (ctrl *postsController) Write(c *fiber.Ctx) error {
	req := new(dto.PostsCreateRequest)
	if;; err := c.BodyParser(req); err != nil {
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
func (ctrl *postsController) List(c *fiber.Ctx) error {
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

	req := &dto.PostsListRequest{
		Page:     pageInt,
		Limit:    limitInt,
		Tag:      tag,
		Username: username,
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsSvc.GetPosts(c, req)
}

// GetPostById  godoc
// @Summary     fetch posts api
// @Description fetch posts by postID
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /posts/:id [post]
func (ctrl *postsController) GetPostById(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	req := &dto.PostsFetchRequest{
		ID: postId,
	}

	return ctrl.postsSvc.GetPost(c, req)
}

// UpdatePostById godoc
// @Summary       update post api
// @Description   update post by id
// @Tags          post
// @Produce       json
// @Success       200
// @Router        /posts [patch]
func (ctrl *postsController) UpdatePostById(c *fiber.Ctx) error {
	req := new(dto.PostsUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsSvc.UpdatePost(c, req)
}

// RemovePostById godoc
// @Summary       remove post api
// @Description   remove post by id
// @Tags          post
// @Produce       json
// @Success       200
// @Router        /posts [delete]
func (ctrl *postsController) RemovePostById(c *fiber.Ctx) error {
	req := new(dto.PostsDeleteRequest)
	if err := c.BodyParser(req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return ctrl.postsSvc.RemovePost(c, req)
}
