package dto

type PostFetchRequest struct {
	ID int `json:"id"`
}

type PostFetchResponse struct {
	PostCreateResponse
	Username string `json:"username"`
}
