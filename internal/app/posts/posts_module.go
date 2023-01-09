package posts

import (
	"github.com/kyh0703/stock-server/internal/types"
)

var PostsModule types.Module

func init() {
	PostsModule.AttachControllers(
		NewPostController(
			NewPostsService(
				NewPostsRepository(),
			),
		),
	)
}
