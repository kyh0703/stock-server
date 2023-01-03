package dto

type PostsUpdateRequest struct {
	ID    int      `json:"id" validate:"required"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type PostsUpdateResponse struct {
	PostsCreateResponse
}
