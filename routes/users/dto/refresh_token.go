package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"require"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken" validate:"require"`
}
