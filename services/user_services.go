package services

import (
	"errors"
	"go-restful-blog-api/v2/auth/token"
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/mappers"
	"go-restful-blog-api/v2/models"
	"go-restful-blog-api/v2/repositories"
	"go-restful-blog-api/v2/utils"

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

	// Hash password sebelum disimpan di database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	// Update password dengan yang sudah di-hash
	user.Password = hashedPassword

	// Panggil repository untuk menyimpan user ke dalam database
	return s.repo.RegisterUser(user)
}

// LoginUser memverifikasi kredensial pengguna
func (s *UserService) LoginUser(UserDTO dto.UserDTO) (dto.UserDTO, string, error) {
	// Mengonversi UserDTO ke dalam model User
	user := mappers.MapToUser(UserDTO)

	// Cari pengguna berdasarkan username melalui repository
	user, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return dto.UserDTO{}, "", errors.New("username tidak ditemukan") // Kembalikan error jika user tidak ditemukan
	}

	// Verifikasi password menggunakan auth/service
	if !utils.ComparePassword(user.Password, UserDTO.Password) {
		return dto.UserDTO{}, "", errors.New("invalid username or password")
	}

	// Generate JWT token jika login berhasil
	token, err := token.GenerateToken(user)
	if err != nil {
		return dto.UserDTO{}, "", err
	}

	UserDTO = mappers.MapToUserDTO(user)
	return UserDTO, token, nil // Kembalikan user jika berhasil login
}
