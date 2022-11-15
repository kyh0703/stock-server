package dto

type PostsListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}

type PostsListResponse struct {
	Posts []PostsFetchResponse `json:"posts"`
}
