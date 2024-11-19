package dto

// UserDTO adalah struktur untuk menerima data user dari client
type UserDTO struct {
	Username string `json:"username" binding:"required"` // Username diperlukan
	Password string `json:"password" binding:"required"` // Password diperlukan
}
