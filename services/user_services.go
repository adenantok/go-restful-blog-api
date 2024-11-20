package services

import (
	"errors"
	"go-restful-blog-api/v2/auth/service"
	"go-restful-blog-api/v2/auth/token"
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
func (s *UserService) RegisterUser(userDTO dto.UserDTO) (models.User, error) {
	// Mengonversi UserDTO ke dalam model User
	user := mappers.MapToUser(userDTO)

	// Validasi apakah username sudah ada di database
	existingUser, err := s.repo.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Jika ada error lain selain record not found
		return models.User{}, err
	}
	if existingUser.ID != 0 { // Mengindikasikan bahwa username sudah ada
		return models.User{}, errors.New("username already exists")
	}

	// Panggil repository untuk menyimpan user ke dalam database
	return s.repo.RegisterUser(user)
}

// LoginUser memverifikasi kredensial pengguna
func (s *UserService) LoginUser(UserDTO dto.UserDTO) (models.User, string, error) {
	// Cari pengguna berdasarkan username melalui repository
	user, err := s.repo.GetUserByUsername(UserDTO.Username)
	if err != nil {
		return models.User{}, "", errors.New("username tidak ditemukan") // Kembalikan error jika user tidak ditemukan
	}

	// Verifikasi password menggunakan auth/service
	if !service.ComparePassword(user.Password, UserDTO.Password) {
		return models.User{}, "", errors.New("invalid username or password")
	}

	// Generate JWT token jika login berhasil
	token, err := token.GenerateToken(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil // Kembalikan user jika berhasil login
}
