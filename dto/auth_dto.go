package dto

type AuthDTO struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
