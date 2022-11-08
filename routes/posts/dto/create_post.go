package dto

type PostCreateRequest struct {
	Title  string   `json:"title" validate:"required"`
	Body   string   `json:"body" validate:"required"`
	Tags   []string `json:"tags" validate:"required"`
	UserID int      `json:"userId"`
}

type PostCreateResponse struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Tags   []string `json:"tags"`
	UserID int      `json:"userId"`
}
