package dto

type PostListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}

type PostListResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	PublishAt string `json:"publishAt"`
	Username  string `json:"username"`
}
