package posts

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/routes/posts/dtos"
	"github.com/kyh0703/stock-server/internal/types"
)

type PostsService struct {
	postsRepo *PostsRepository
}

func NewPostsService(postsRepo *PostsRepository) *PostsService {
	return &PostsService{
		postsRepo: postsRepo,
	}
}

func (svc *PostsService) Create(c *fiber.Ctx, dto *dtos.CreatePostDto) error {
	post, err := svc.postsRepo.Save(
		c.Context(),
		dto.Title,
		dto.Body,
		dto.Tags,
		dto.UserID,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	return c.Status(fiber.StatusOK).JSON(&dtos.PostsDto{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		Tags:      post.Tags,
		PublishAt: post.PublishAt.String(),
		UserID:    post.Edges.User.ID,
	})
}

func (svc *PostsService) FindOne(c *fiber.Ctx, id int) error {
	post, err := svc.postsRepo.FetchOneWithUser(c.Context(), id)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	var res dtos.PostsDto
	res.ID = post.ID
	res.Title = post.Title
	res.Body = post.Body
	res.Tags = post.Tags
	res.PublishAt = post.PublishAt.String()
	res.UserID = post.Edges.User.ID

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *PostsService) GetPosts(c *fiber.Ctx, dto *dtos.FindPostsDto) error {
	_, err := svc.postsRepo.FetchPostsWithTagOrUser(
		c.Context(),
		dto.Tag,
		dto.Username,
		dto.Page,
		dto.Limit,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	count, err := svc.postsRepo.CountByNameOrTag(
		c.Context(),
		dto.Tag,
		dto.Username,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	lastPage := strconv.Itoa(
		int(math.Ceil(float64(count) / float64(dto.Limit))),
	)
	c.Response().Header.Set("last-page", lastPage)

	var res dtos.FindPostsDto
	// for _, v := range posts {
	// 	var post dtos.PostsFetchResponse
	// 	post.ID = v.ID
	// 	post.Title = v.Title
	// 	post.Body = v.Body
	// 	post.Tags = v.Tags
	// 	post.PublishAt = v.PublishAt.String()
	// 	post.UserID = v.Edges.User.ID
	// 	post.Username = v.Edges.User.Username
	// 	res.Posts = append(res.Posts, post)
	// }

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *PostsService) Update(c *fiber.Ctx, id int, dto *dtos.CreatePostDto) error {
	if err := svc.CheckOwnPost(c, id); err != nil {
		return err
	}

	post, err := svc.postsRepo.Update(
		c.Context(),
		id,
		dto.Title,
		dto.Body,
		dto.Tags,
	)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	var res dtos.PostsDto
	res.ID = post.ID
	res.Title = post.Title
	res.Body = post.Body
	res.Tags = post.Tags
	res.PublishAt = post.PublishAt.String()

	return c.Status(fiber.StatusOK).JSON(res)
}

func (svc *PostsService) RemovePost(c *fiber.Ctx, id int) error {
	if err := svc.CheckOwnPost(c, id); err != nil {
		return err
	}

	if err := svc.postsRepo.Remove(c.Context(), id); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (svc *PostsService) CheckOwnPost(c *fiber.Ctx, id int) error {
	userID, ok := c.Context().UserValue(types.ContextKeyUserID).(int)
	if !ok {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	post, err := svc.postsRepo.FetchOneWithUser(c.Context(), id)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	if userID != post.Edges.User.ID {
		return c.App().ErrorHandler(c, types.ErrUserUnauthorized)
	}

	return nil
}
