package posts

import (
	"context"
	"testing"

	"github.com/kyh0703/stock-server/database"
	"github.com/stretchr/testify/assert"
)

func TestFetchPost(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	client, err := database.ConnectDb(ctx)
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

	client, err := database.ConnectDb(ctx)
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

	client, err := database.ConnectDb(ctx)
	assert.NoError(t, err)
	defer client.Close()

	var repo PostsRepository
	posts, err := repo.FetchPostsWithTagOrUser(ctx, "", "", 1, 10)
	assert.NoError(t, err)
	t.Log(posts)
}
