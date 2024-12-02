package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	PostID    int       `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    int       `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relasi dengan Post
	Post Post `gorm:"foreignKey:PostID"`

	// Relasi dengan User
	User User `gorm:"foreignKey:UserID"`
}

func MigrateComment(db *gorm.DB) error {
	return db.AutoMigrate(&Comment{})
}
