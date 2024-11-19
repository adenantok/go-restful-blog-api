package models

import (
	"gorm.io/gorm"
)

// User mendefinisikan struktur model untuk pengguna.
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"` // ID adalah primary key dan auto increment
	Username string `gorm:"unique;not null"`          // Username harus unik dan tidak boleh null
	Password string `gorm:"not null"`                 // Password tidak boleh null
}

// MigrateUser menjalankan AutoMigrate untuk memastikan tabel user sesuai dengan model
func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
