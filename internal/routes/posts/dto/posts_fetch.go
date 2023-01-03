package dto

type PostsFetchRequest struct {
	ID int `json:"id"`
}

type PostsFetchResponse struct {
	PostsCreateResponse
	Username string `json:"username"`
}
