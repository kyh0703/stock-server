package dtos

import "github.com/kyh0703/stock-server/ent"

type PostDto struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	PublishAt string   `json:"publishAt"`
	UserID    int      `json:"userID"`
}

func (dto *PostDto) Serialize(post *ent.Post) *PostDto {
	dto.ID = post.ID
	dto.Title = post.Title
	dto.Body = post.Body
	dto.Tags = post.Tags
	dto.PublishAt = post.PublishAt.String()
	if post.Edges.User != nil {
		dto.UserID = post.Edges.User.ID
	}
	return dto
}
