package dto

type CreateUserDTO struct {
	Email           string `json:"email" validate:"required"`
	Username        string `json:"username" validate:"required,min=2,max=20"`
	Password        string `json:"password" validate:"required,min=6,max=20"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=6,max=20"`
}
