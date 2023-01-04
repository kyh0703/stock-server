package posts

import (
	"context"
	"testing"
	"time"

	"github.com/kyh0703/stock-server/configs"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPostRepository struct {
	suite.Suite
	client *ent.Client
}

// suite 전체 테스트 실행 전에 한번 실행
func (test *TestPostRepository) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	client, err := database.ConnectDatabase(
		ctx,
		configs.Env.DBType,
		configs.Env.DBUrl,
	)
	assert.NoError(test.T(), err)
	test.client = client
}

// 각 테스트 전에 실행
func (test *TestPostRepository) SetupTest() {
}

func (test *TestPostRepository) TestFetchPost(t *testing.T) {
	var (
		repo   PostsRepository
		postID int
	)
	post, err := repo.FetchOne(context.TODO(), postID)
	assert.NoError(t, err)
	t.Log(post)
}

func (test *TestPostRepository) TestFetchPostJoin() {
	var (
		repo   PostsRepository
		postID int
	)
	post, err := repo.FetchOneWithUser(context.TODO(), postID)
	assert.NoError(test.T(), err)
	test.T().Log(post)
}

func (test *TestPostRepository) TestFetchPostsWithTagOrUser() {
	var repo PostsRepository
	posts, err := repo.FetchPostsWithTagOrUser(context.TODO(), "", "", 1, 10)
	assert.NoError(test.T(), err)
	test.T().Log(posts)
}

// 각 테스트가 종료 시
func (test *TestPostRepository) TerDownTest() {
}

// 전체 테스트가 종료 시
func (test *TestPostRepository) TearDownSuite() {
	defer test.client.Close()
}
