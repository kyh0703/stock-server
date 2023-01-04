package dtos

type UsersRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"require"`
}
