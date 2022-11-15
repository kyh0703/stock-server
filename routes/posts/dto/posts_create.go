package dto

type PostsCreateRequest struct {
	Title  string   `json:"title" validate:"required"`
	Body   string   `json:"body" validate:"required"`
	Tags   []string `json:"tags" validate:"required"`
	UserID int      `json:"userId" validate:"required"`
}

type PostsCreateResponse struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	PublishAt string   `json:"publishAt"`
	UserID    int      `json:"userId"`
}
