package posts

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/routes/posts/dto"
	"github.com/kyh0703/stock-server/types"
)

type postsService struct {
	postRepo PostsRepository
}

func (svc *postsService) SavePost(c *fiber.Ctx, req *dto.PostCreateRequest) error {
	post, err := svc.postRepo.Insert(
		c.Context(),
		req.Title,
		req.Body,
		req.Tags,
		req.UserID,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	res := &dto.PostCreateResponse{
		ID:     post.ID,
		Title:  post.Title,
		Body:   post.Body,
		Tags:   post.Tags,
		UserID: req.UserID,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *postsService) GetPost(c *fiber.Ctx, id int) error {
	post, err := svc.postRepo.FetchOne(c.Context(), id)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func (svc *postsService) GetPosts(c *fiber.Ctx, req *dto.PostListRequest) error {
	posts, err := svc.postRepo.FetchPostsWithTagOrUser(
		c.Context(),
		req.Tag,
		req.Username,
		req.Page,
		req.Limit,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	count, err := svc.postRepo.CountByNameOrTag(
		c.Context(),
		req.Tag,
		req.Username,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	c.Response().Header.Set("last-page", strconv.Itoa(int(math.Ceil(float64(count/10)))))

	return c.Status(fiber.StatusOK).JSON(posts)
}
