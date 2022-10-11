package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/api/middleware"
	"github.com/kyh0703/stock-server/ent"
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
	// db, _ := c.Keys["database"].(*ent.Client)
	// pagination
	// var (
	// 	page  = c.DefaultQuery("page", "1")
	// 	tag   = c.Query("tag")
	// 	email = c.Query("email")
	// )
	// db.Post.Query().
	// 	Where(post.Or(
	// 		post.tagoo
	// 	))
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
