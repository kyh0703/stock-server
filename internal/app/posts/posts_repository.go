package posts

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/user"
	"github.com/kyh0703/stock-server/pkg/database"
)

type PostsRepository struct{}

func NewPostsRepository() *PostsRepository {
	return &PostsRepository{}
}

func (repo *PostsRepository) Save(ctx context.Context, title, body string, tags []string, userID int) (*ent.Post, error) {
	return database.Ent.Post.
		Create().
		SetTitle(title).
		SetBody(body).
		SetTags(tags).
		SetUserID(userID).
		Save(ctx)
}

func (repo *PostsRepository) Update(ctx context.Context, id int, title, body string, tags []string) (*ent.Post, error) {
	return database.Ent.Post.
		UpdateOneID(id).
		SetTitle(title).
		SetBody(body).
		SetTags(tags).
		Save(ctx)
}

func (repo *PostsRepository) Remove(ctx context.Context, id int) error {
	return database.Ent.Post.
		DeleteOneID(id).
		Exec(ctx)
}

func (repo *PostsRepository) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	return database.Ent.Post.
		Query().
		Where(post.ID(id)).
		WithUser().
		Only(ctx)
}

func (repo *PostsRepository) PagePosts(ctx context.Context, tag, username string, page, limit int) (ent.Posts, error) {
	return database.Ent.Post.
		Query().
		Select().
		Limit(limit).
		Offset((page - 1) * limit).
		WithUser().
		Where(
			post.Or(
				post.HasUserWith(user.UsernameContains(username)),
				func(s *sql.Selector) {
					s.Where(sqljson.StringContains(post.FieldTags, tag))
				},
			),
		).
		Order(ent.Desc(post.FieldPublishAt)).
		All(ctx)
}
