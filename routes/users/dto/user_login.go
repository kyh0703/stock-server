package dto

type UserLoginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=10"`
}
