package dto

type UsersRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"require"`
}

type UsersRefreshResponse struct {
	AccessToken string `json:"accessToken" validate:"require"`
}
