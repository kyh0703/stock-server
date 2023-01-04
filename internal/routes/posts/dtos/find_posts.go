package dtos

type FindPostsDto struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}
