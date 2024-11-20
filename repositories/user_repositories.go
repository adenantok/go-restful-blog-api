package repositories

import (
	"errors" // Mengimport package untuk hashing password
	"go-restful-blog-api/v2/models"
	"go-restful-blog-api/v2/utils"

	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak repositori untuk user
type UserRepository interface {
	RegisterUser(user models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

// userRepository adalah implementasi dari UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// RegisterUser menerima model User dan menyimpannya ke dalam database
func (repo *userRepository) RegisterUser(user models.User) (models.User, error) {
	// Hash password sebelum disimpan di database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	// Update password dengan yang sudah di-hash
	user.Password = hashedPassword

	// Simpan user ke database
	if err := repo.db.Create(&user).Error; err != nil {
		// Periksa apakah username sudah ada
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return models.User{}, errors.New("username already exists")
		}
		return models.User{}, err
	}

	return user, nil
}

// GetUserByUsername mencari user berdasarkan username
func (repo *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
