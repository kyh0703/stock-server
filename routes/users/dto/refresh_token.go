package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"require"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token" validate:"require"`
}
