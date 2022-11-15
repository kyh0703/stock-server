package dto

import "github.com/kyh0703/stock-server/ent"

type PostsListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}

type PostsListResponse struct {
	Posts []PostsFetchResponse `json:"posts"`
}

func (res *PostsListResponse) SetEntity(posts []*ent.Posts) {
}
