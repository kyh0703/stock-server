package posts

import (
	"context"
	"testing"

	"github.com/kyh0703/stock-server/configs"
	"github.com/kyh0703/stock-server/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPostRepository struct {
	suite.Suite
}

func (suite *TestPostRepository) SetupTest() {
}

func TestFetchPost(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	client, err := database.ConnectDatabase(
		ctx,
		configs.Env.DBType,
		configs.Env.DBUrl,
	)
	assert.NoError(t, err)
	defer client.Close()

	var (
		repo   PostsRepository
		postID int
	)
	post, err := repo.FetchOne(ctx, postID)
	assert.NoError(t, err)
	t.Log(post)
}

func TestFetchPostJoin(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	client, err := database.ConnectDatabase(
		ctx,
		configs.Env.DBType,
		configs.Env.DBUrl,
	)
	assert.NoError(t, err)
	defer client.Close()

	var (
		repo   PostsRepository
		postID int
	)

	post, err := repo.FetchOneWithUser(ctx, postID)
	assert.NoError(t, err)
	t.Log(post)
}

func TestFetchPostsWithTagOrUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	client, err := database.ConnectDatabase(ctx,
		configs.Env.DBType,
		configs.Env.DBUrl,
	)
	assert.NoError(t, err)
	defer client.Close()

	var repo PostsRepository
	posts, err := repo.FetchPostsWithTagOrUser(ctx, "", "", 1, 10)
	assert.NoError(t, err)
	t.Log(posts)
}
