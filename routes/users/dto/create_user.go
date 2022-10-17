package dto

type CreateUserDTO struct {
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=3,max=10"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=3,max=10"`
	Name            string `json:"name" validate:"required"`
}
