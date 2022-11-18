package posts

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/routes/posts/dto"
	"github.com/kyh0703/stock-server/types"
)

type postsService struct {
	postsRepo PostsRepository
}

func (svc *postsService) SavePost(c *fiber.Ctx, req *dto.PostsCreateRequest) error {
	post, err := svc.postsRepo.Insert(
		c.Context(),
		req.Title,
		req.Body,
		req.Tags,
		req.UserID,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	res := new(dto.PostsCreateResponse)
	res.ID = post.ID
	res.Title = post.Title
	res.Body = post.Body
	res.Tags = post.Tags
	res.PublishAt = post.PublishAt.String()
	res.UserID = req.UserID

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *postsService) GetPost(c *fiber.Ctx, req *dto.PostsFetchRequest) error {
	post, err := svc.postsRepo.FetchOneWithUser(c.Context(), req.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	var res dto.PostsFetchResponse
	res.ID = post.ID
	res.Title = post.Title
	res.Body = post.Body
	res.Tags = post.Tags
	res.PublishAt = post.PublishAt.String()
	res.UserID = post.Edges.User.ID
	res.Username = post.Edges.User.Username

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *postsService) GetPosts(c *fiber.Ctx, req *dto.PostsListRequest) error {
	posts, err := svc.postsRepo.FetchPostsWithTagOrUser(
		c.Context(),
		req.Tag,
		req.Username,
		req.Page,
		req.Limit,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	count, err := svc.postsRepo.CountByNameOrTag(
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

	var res dto.PostsListResponse
	for _, v := range posts {
		var post dto.PostsFetchResponse
		post.ID = v.ID
		post.Title = v.Title
		post.Body = v.Body
		post.Tags = v.Tags
		post.PublishAt = v.PublishAt.String()
		post.UserID = v.Edges.User.ID
		post.Username = v.Edges.User.Username
		res.Posts = append(res.Posts, post)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *postsService) UpdatePost(c *fiber.Ctx, req *dto.PostsUpdateRequest) error {
	if err := svc.CheckOwnPost(c, req.ID); err != nil {
		return err
	}

	post, err := svc.postsRepo.UpdateById(c.Context(), req.ID, req.Title, req.Body, req.Tags)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	var res dto.PostsUpdateResponse
	res.ID = post.ID
	res.Title = post.Title
	res.Body = post.Body
	res.Tags = post.Tags
	res.PublishAt = post.PublishAt.String()

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *postsService) RemovePost(c *fiber.Ctx, req *dto.PostsDeleteRequest) error {
	if err := svc.CheckOwnPost(c, req.ID); err != nil {
		return err
	}

	if err := svc.postsRepo.DeleteById(c.Context(), req.ID); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (svc *postsService) CheckOwnPost(c *fiber.Ctx, postId int) error {
	userId, ok := c.UserContext().Value("user_id").(int)
	if !ok {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	post, err := svc.postsRepo.FetchOneWithUser(c.Context(), postId)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	if userId != post.Edges.User.ID {
		return c.App().ErrorHandler(c, types.ErrUserUnauthorized)
	}

	return nil
}
