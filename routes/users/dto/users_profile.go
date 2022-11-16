package dto

type UsersProfileRequest struct {
	ID int `json:"id" validate:"require"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
