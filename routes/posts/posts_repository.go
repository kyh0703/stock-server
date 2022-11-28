package posts

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/posts"
	"github.com/kyh0703/stock-server/ent/users"
)

type PostsRepository struct{}

func (repo *PostsRepository) Insert(ctx context.Context, title, body string, tags []string, userID int) (*ent.Posts, error) {
	return database.Ent.Posts.
		Create().
		SetTitle(title).
		SetBody(body).
		SetTags(tags).
		SetUserID(userID).
		Save(ctx)
}

func (repo *PostsRepository) UpdateById(ctx context.Context, id int, title, body string, tags []string) (*ent.Posts, error) {
	return database.Ent.Posts.
		UpdateOneID(id).
		SetTitle(title).
		SetBody(body).
		SetTags(tags).
		Save(ctx)
}

func (repo *PostsRepository) DeleteById(ctx context.Context, id int) error {
	return database.Ent.Posts.
		DeleteOneID(id).
		Exec(ctx)
}

func (repo *PostsRepository) FetchOne(ctx context.Context, id int) (*ent.Posts, error) {
	return database.Ent.Posts.
		Query().
		Where(posts.ID(id)).
		Only(ctx)
}

func (repo *PostsRepository) FetchOneWithUser(ctx context.Context, id int) (*ent.Posts, error) {
	return database.Ent.Posts.
		Query().
		Where(posts.ID(id)).
		WithUser().
		Only(ctx)
}

func (repo *PostsRepository) FetchPostsWithTagOrUser(ctx context.Context, tag, username string, page, limit int) ([]*ent.Posts, error) {
	return database.Ent.Posts.
		Query().
		Select().
		Limit(limit).
		Offset((page - 1) * limit).
		WithUser().
		Where(
			posts.Or(
				posts.HasUserWith(users.UsernameContains(username)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(posts.FieldTags, tag))
				},
			),
		).
		Order(ent.Desc(posts.FieldPublishAt)).
		All(ctx)
}

func (repo *PostsRepository) CountByNameOrTag(ctx context.Context, tag, name string) (int, error) {
	return database.Ent.Posts.
		Query().
		Where(
			posts.Or(
				posts.HasUserWith(users.UsernameContains(name)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(posts.FieldTags, tag))
				},
			),
		).
		Count(ctx)
}
