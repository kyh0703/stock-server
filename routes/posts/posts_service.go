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

func (svc *postsService) GetPost(c *fiber.Ctx, req *dto.PostFetchRequest) error {
	post, err := svc.postRepo.FetchOneWithUser(c.Context(), req.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	res := &dto.PostFetchResponse{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		PublishAt: post.PublishAt.String(),
		UserID:    post.Edges.User.ID,
		Email:     post.Edges.User.Email,
	}

	return c.Status(fiber.StatusOK).JSON(res)
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

	lastPage := strconv.Itoa(
		int(math.Ceil(float64(count) / float64(req.Limit))),
	)
	c.Response().Header.Set("last-page", lastPage)
	return c.Status(fiber.StatusOK).JSON(posts)
}

func (svc *postsService) RemovePost(c *fiber.Ctx, id int) error {
	if err := svc.postRepo.DeleteById(c.Context(), id); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
