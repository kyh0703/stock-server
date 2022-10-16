package v1

import (
	"math"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/predicate"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/middleware"
)

type postController struct {
	path   string
	router fiber.Router
}

func NewPostController(router fiber.Router) *postController {
	return &postController{
		path:   "/post",
		router: router,
	}
}

func (ctrl *postController) Index() fiber.Router {
	r := ctrl.router.Group(ctrl.path)
	r.Get("/", ctrl.List)
	r.Get("/:id", ctrl.GetPostById)
	r.Post("/write", middleware.TokenAuth(), ctrl.Write)
	return r
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) Write(c *fiber.Ctx) error {
	req := struct {
		Title string   `json:"title" validate:"required"`
		Body  string   `json:"body" validate:"required"`
		Tags  []string `json:"tags" validate:"required"`
	}{}
	// body parser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// validate request message
	if err := validator.New().StructCtx(c.Context(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	// save the database
	post, err := database.Ent().Post.
		Create().
		SetTitle(req.Title).
		SetBody(req.Body).
		SetTags(req.Tags).
		SetUserID(userID).
		Save(c.Context())
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
		page     = c.Query("page", "1")
		tag      = c.Query("tag")
		username = c.Query("username")
	)
	// parse page
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// make query
	var query predicate.Post
	if tag != "" {
		query = func(s *sql.Selector) {
			s.Where(sqljson.StringContains(post.FieldTags, tag))
		}
	}
	if username != "" {
		query = post.HasUserWith(user.UsernameContains(username))
	}
	// select data
	posts, err := database.Ent().Post.
		Query().
		Limit(10).
		Offset((pageInt - 1) * 10).
		Where(query).
		All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// get post count
	postCount, err := database.Ent().Post.
		Query().
		Where(query).
		Count(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	c.Response().Header.Set("last-page", strconv.Itoa(int(math.Ceil(float64(postCount/10)))))
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
	param := c.Params("id")
	if param == "" {
		return fiber.ErrBadRequest
	}
	// strconv
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// get post data
	post, err := database.Ent().Post.
		Query().
		Where(post.ID(id)).
		Only(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}
