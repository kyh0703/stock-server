package dto

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=10"`
}

type UserLoginResponse struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	AccessToken        string `json:"accessToken"`
	AccessTokenExpires int64  `json:"accessTokenExpire"`
}
