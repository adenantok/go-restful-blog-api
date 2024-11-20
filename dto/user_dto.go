package dto

// UserDTO adalah struktur untuk menerima data user dari client
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"` // Username diperlukan
	Password string `json:"password" binding:"required"` // Password diperlukan
}
