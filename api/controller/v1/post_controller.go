package v1

import (
	"math"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/api/middleware"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/predicate"
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
	route.POST("/write", middleware.CheckLoggedIn(), ctrl.Write)
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
func (ctrl *postController) List(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// pagination
	var (
		page = c.DefaultQuery("page", "1")
		tag  = c.Query("tag")
		_    = c.Query("email")
	)
	// parse page
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// make query
	var query []predicate.Post
	query = append(query, func(s *sql.Selector) {
		s.Where(sqljson.StringContains(post.FieldTags, tag))
	})
	// select data
	posts, err := db.Post.
		Query().
		Limit(10).
		Offset((pageInt - 1) * 10).
		Where(post.Or(query...)).
		All(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// get post count
	postCount, err := db.Post.Query().Where(post.Or(query...)).Count(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Writer.Header().Set("Last-Page", strconv.Itoa(int(math.Ceil(float64(postCount/10)))))
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
func (ctrl *postController) Write(c *gin.Context) {
	db, _ := c.Keys["database"].(*ent.Client)
	// validator
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
	post, err := db.Post.Create().
		SetTitle(req.Title).
		SetBody(req.Body).
		SetTags(req.Tags).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, post)
}
