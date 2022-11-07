package dto

type PostListResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	PublishAt string `json:"publishAt"`
	Username  string `json:"username"`
}
