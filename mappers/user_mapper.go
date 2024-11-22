package mappers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/models"
)

// MapToUser mengonversi UserDTO ke User model
func MapToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		Username: userDTO.Username,
		Password: userDTO.Password, // Password akan dihash sebelum disimpan di database
	}
}

// MapToUserDTO mengonversi User model ke UserDTO
func MapToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		//Password: user.Password, // Pastikan tidak ada field ID
	}
}
