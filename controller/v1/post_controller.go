package v1

import (
	"math"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/predicate"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/middleware"
)

type postController struct {
	path string
	rg   *gin.RouterGroup
}

func NewPostController(rg *gin.RouterGroup) *postController {
	return &postController{
		path: "/post",
		rg:   rg,
	}
}

func (ctrl *postController) Index() *gin.RouterGroup {
	route := ctrl.rg.Group(ctrl.path)
	route.GET("/", ctrl.List)
	route.GET("/:id", ctrl.GetPostById)
	route.POST("/write", middleware.TokenAuth(), ctrl.Write)
	return route
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Param       username string
// @Param       password string
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) Write(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// validator
	userID, err := strconv.Atoi(c.Request.Header.Get("x-request-id"))
	if err != nil || userID == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	req := struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags" binding:"required"`
	}{}
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// save the database
	post, err := db.Post.
		Create().
		SetTitle(req.Title).
		SetBody(req.Body).
		SetTags(req.Tags).
		SetUserID(userID).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Param       username string
// @Param       password string
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) List(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// get query data
	var (
		page     = c.DefaultQuery("page", "1")
		tag      = c.Query("tag")
		username = c.Query("username")
	)
	// parse page
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
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
	posts, err := db.Post.
		Query().
		Limit(10).
		Offset((pageInt - 1) * 10).
		Where(query).
		All(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// get post count
	postCount, err := db.Post.Query().Where(query).Count(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Writer.Header().Set("last-page", strconv.Itoa(int(math.Ceil(float64(postCount/10)))))
	c.JSON(http.StatusOK, posts)
}

// Register     godoc
// @Summary     register auth info
// @Description register stock api
// @Tags        auth
// @Produce     json
// @Param       username string
// @Param       password string
// @Success     200
// @Router      /auth/register [post]
func (ctrl *postController) GetPostById(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// validate
	param := c.Param("id")
	if param == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// strconv
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// get post data
	post, err := db.Post.
		Query().
		Where(post.ID(id)).
		Only(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, post)
}
