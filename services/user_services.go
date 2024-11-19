package services

import (
	"errors"
	auth "go-restful-blog-api/v2/auth/service"
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/mappers"
	"go-restful-blog-api/v2/models"
	"go-restful-blog-api/v2/repositories"

	"gorm.io/gorm"
)

// userService adalah struct yang akan menangani logika bisnis pengguna
type UserService struct {
	repo repositories.UserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// RegisterUser menerima data UserDTO, memvalidasi dan menyimpannya ke database
func (service *UserService) RegisterUser(userDTO dto.UserDTO) (models.User, error) {
	// Mengonversi UserDTO ke dalam model User
	user := mappers.MapToUser(userDTO)

	// Hash password sebelum disimpan di database
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	// Update password dengan yang sudah di-hash
	user.Password = hashedPassword

	// Validasi apakah username sudah ada di database
	existingUser, err := service.repo.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Jika ada error lain selain record not found
		return models.User{}, err
	}
	if existingUser.ID != 0 { // Mengindikasikan bahwa username sudah ada
		return models.User{}, errors.New("username already exists")
	}

	// Panggil repository untuk menyimpan user ke dalam database
	return service.repo.RegisterUser(user)
}
