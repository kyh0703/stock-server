package dto

type UserProfileRequest struct {
	ID int `json:"id"`
}

type UserProfileResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
