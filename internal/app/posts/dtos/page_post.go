package dtos

type PagePostsDto struct {
	Page     int    `json:"page" validate:"default=1"`
	Limit    int    `json:"limit" validate:"default=10"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
}
