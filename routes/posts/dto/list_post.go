package dto

type PostListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}

type PostListResponse struct {
	PostFetchResponse
}
