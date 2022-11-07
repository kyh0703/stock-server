package posts

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/user"
)

type PostsRepository struct{}

func (repository *PostsRepository) Save(ctx context.Context, post *ent.Post, userId int) (*ent.Post, error) {
	return database.Ent.Post.
		Create().
		SetTitle(post.Title).
		SetBody(post.Body).
		SetTags(post.Tags).
		SetUserID(userId).
		Save(ctx)
}

func (repository *PostsRepository) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	return database.Ent.Post.
		Query().
		Select(
			post.FieldID,
			post.FieldTitle,
			post.FieldBody,
			post.FieldTags,
		).
		Where(post.ID(id)).
		Only(ctx)
}

func (repository *PostsRepository) FetchPostsWithTagOrUser(ctx context.Context, tag, username string, page, limit int) ([]*ent.Post, error) {
	return database.Ent.Debug().Post.
		Query().
		Select(post.FieldID, post.FieldTitle, post.FieldBody, post.FieldTags).
		Limit(limit).
		Offset((page - 1) * limit).
		Where(
			post.And(
				post.HasUserWith(user.UsernameContains(username)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(post.FieldTags, tag))
				},
			),
		).
		All(ctx)
}

func (repository *PostsRepository) CountByNameOrTag(ctx context.Context, tag, name string) (int, error) {
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
