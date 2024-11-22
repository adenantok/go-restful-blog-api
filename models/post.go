package models

import (
	"time"

	"gorm.io/gorm"
)

// Post represents the structure of a blog post
type Post struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relasi dengan User
	User User `gorm:"foreignKey:UserID"`
}

func MigratePost(db *gorm.DB) error {
	return db.AutoMigrate(&Post{})
}
