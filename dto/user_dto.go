package dto

// UserDTO adalah struktur untuk menerima data user dari client
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"` // Username diperlukan
	Password string `json:"-"`
}

// UserCreateDTO digunakan untuk menerima data input (termasuk password)
type UserCreateDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` // Validasi password saat input
}
