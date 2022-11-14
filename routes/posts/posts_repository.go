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

func (repo *PostsRepository) Insert(ctx context.Context, title, body string, tags []string, userId int) (*ent.Post, error) {
	return database.Ent.Debug().Post.
		Create().
		SetTitle(title).
		SetBody(body).
		SetTags(tags).
		SetUserID(userId).
		Save(ctx)
}

func (repo *PostsRepository) DeleteById(ctx context.Context, id int) error {
	return database.Ent.Debug().Post.
		DeleteOneID(id).
		Exec(ctx)
}

func (repo *PostsRepository) FetchOne(ctx context.Context, id int) (*ent.Post, error) {
	return database.Ent.Debug().Post.
		Query().
		Where(post.ID(id)).
		Only(ctx)
}

func (repo *PostsRepository) FetchOneWithUser(ctx context.Context, id int) (*ent.Post, error) {
	return database.Ent.Debug().Post.
		Query().
		Where(post.ID(id)).
		WithUser().
		Only(ctx)
}

func (repo *PostsRepository) FetchPostsWithTagOrUser(ctx context.Context, tag, username string, page, limit int) ([]*ent.Post, error) {
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

func (repo *PostsRepository) CountByNameOrTag(ctx context.Context, tag, name string) (int, error) {
	return database.Ent.Debug().Post.
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
