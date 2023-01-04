package dtos

type PostsDto struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	PublishAt string   `json:"publishAt"`
	UserID    int      `json:"userID"`
}
