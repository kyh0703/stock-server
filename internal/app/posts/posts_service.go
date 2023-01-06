package posts

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/app/posts/dtos"
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

func (svc *PostsService) Create(c *fiber.Ctx, dto *dtos.CreatePostDto) (*dtos.PostDto, error) {
	post, err := svc.postsRepo.Save(
		c.Context(),
		dto.Title,
		dto.Body,
		dto.Tags,
		dto.UserID,
	)
	if err != nil {
		return nil, c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	return new(dtos.PostDto).Serialize(post), nil
}

func (svc *PostsService) FindOne(c *fiber.Ctx, id int) (*dtos.PostDto, error) {
	post, err := svc.postsRepo.FindOne(c.Context(), id)
	if err != nil {
		return nil, c.App().ErrorHandler(c, types.ErrPostNotExist)
	}

	return new(dtos.PostDto).Serialize(post), nil
}

func (svc *PostsService) PagePosts(c *fiber.Ctx, dto *dtos.PagePostsDto) ([]*dtos.PostDto, error) {
	posts, err := svc.postsRepo.PagePosts(
		c.Context(),
		dto.Tag,
		dto.Username,
		dto.Page,
		dto.Limit,
	)
	if err != nil {
		return nil, c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	c.Response().Header.Set("last-page",
		strconv.Itoa(
			int(math.Ceil(float64(len(posts))/float64(dto.Limit))),
		),
	)

	var res []*dtos.PostDto
	for _, v := range posts {
		res = append(res, new(dtos.PostDto).Serialize(v))
	}

	return res, nil
}

func (svc *PostsService) Update(c *fiber.Ctx, id int, dto *dtos.CreatePostDto) (*dtos.PostDto, error) {
	if err := svc.CheckOwnPost(c, id); err != nil {
		return nil, err
	}

	post, err := svc.postsRepo.Update(
		c.Context(),
		id,
		dto.Title,
		dto.Body,
		dto.Tags,
	)
	if err != nil {
		return nil, c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	return new(dtos.PostDto).Serialize(post), nil
}

func (svc *PostsService) RemovePost(c *fiber.Ctx, id int) error {
	if err := svc.CheckOwnPost(c, id); err != nil {
		return err
	}

	if err := svc.postsRepo.Remove(c.Context(), id); err != nil {
		return c.App().ErrorHandler(c, types.ErrInvalidParameter)
	}

	return nil
}

func (svc *PostsService) CheckOwnPost(c *fiber.Ctx, id int) error {
	userID, ok := c.Context().UserValue(types.ContextKeyUserID).(int)
	if !ok {
		return c.App().ErrorHandler(c, types.ErrUnauthorized)
	}

	post, err := svc.postsRepo.FindOne(c.Context(), id)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	if userID != post.Edges.User.ID {
		return c.App().ErrorHandler(c, types.ErrUserUnauthorized)
	}

	return nil
}
