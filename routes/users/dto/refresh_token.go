package dto

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"require"`
}
