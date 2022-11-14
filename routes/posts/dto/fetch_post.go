package dto

type PostFetchRequest struct {
	ID int `json:"id"`
}

type PostFetchResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	PublishAt string `json:"publishAt"`
	UserID    int    `json:"userId"`
	Email     string `json:"email"`
}
