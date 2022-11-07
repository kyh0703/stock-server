package posts

import (
	"github.com/gofiber/fiber/v2"
	postsdto "github.com/kyh0703/stock-server/routes/posts/dto"
)

type postsService struct {
	postsRepository PostsRepository
}

func (svc *postsService) RegisterPost(c *fiber.Ctx, dto postsdto.PostCreateRequest) error {
	return nil
}

func (svc *postsService) FetchPost(c *fiber.Ctx, id int) error {
	return nil
}

func (svc *postsService) FetchPosts(c *fiber.Ctx, tag, name string, page, limit int) error {
	return nil
}
