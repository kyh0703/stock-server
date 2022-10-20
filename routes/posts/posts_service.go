package posts

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/user"
	postsdto "github.com/kyh0703/stock-server/routes/posts/dto"
)

type postsService struct{}

func (svc *postsService) SavePost(ctx context.Context, dto postsdto.CreatePostDTO) (*ent.Post, error) {
	return database.Ent.Post.
		Create().
		SetTitle(dto.Title).
		SetBody(dto.Body).
		SetTags(dto.Tags).
		SetUserID(dto.UserID).
		Save(ctx)
}

func (svc *postsService) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	return database.Ent.Post.
		Query().
		Select(post.FieldID, post.FieldTitle, post.FieldBody, post.FieldTags, post.FieldPublishAt).
		Where(post.ID(id)).
		Only(ctx)
}

func (svc *postsService) FindPagesByNameOrTag(ctx context.Context, tag, name string, page, limit int) ([]*ent.Post, error) {
	return database.Ent.Post.
		Query().
		Select(post.FieldID, post.FieldTitle, post.FieldBody, post.FieldTags, post.FieldPublishAt).
		Limit(limit).
		Offset((page - 1) * limit).
		Where(
			post.And(
				post.HasUserWith(user.UsernameContains(name)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(post.FieldTags, tag))
				},
			),
		).
		All(ctx)
}

func (svc *postsService) GetCountByNameOrTag(ctx context.Context, tag, name string) (int, error) {
	return database.Ent.Post.
		Query().
		Where(
			post.And(
				post.HasUserWith(user.UsernameContains(name)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(post.FieldTags, tag))
				},
			),
		).
		Count(ctx)
}
