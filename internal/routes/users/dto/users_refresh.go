package dto

type UsersRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"require"`
}
