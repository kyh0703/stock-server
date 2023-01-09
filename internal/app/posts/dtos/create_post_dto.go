package dtos

type CreatePostDto struct {
	Title  string   `json:"title" validate:"required"`
	Body   string   `json:"body" validate:"required"`
	Tags   []string `json:"tags" validate:"required"`
	UserID int      `json:"userID" validate:"required"`
}