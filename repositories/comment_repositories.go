package repositories

import (
	"go-restful-blog-api/v2/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) (models.Comment, error)
	CheckPostExists(postID int) (bool, error)
	//GetCommentsByPostID(postID int) ([]models.Comment, error)
	//GetCommentByID(ID int) (models.Comment, error)
	//UpdateComment(comment *models.Comment) (models.Comment, error)
	//DeleteComment(ID int) error
}

type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository menginisialisasi CommentRepository baru
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

// CheckPostExists memeriksa apakah postingan dengan ID tertentu ada di database
func (repo *commentRepository) CheckPostExists(postID int) (bool, error) {
	var count int64
	// Menghitung jumlah postingan dengan ID tertentu
	err := repo.db.Model(&models.Post{}).Where("id = ?", postID).Count(&count).Error
	if err != nil {
		return false, err // Jika terjadi error saat query
	}
	return count > 0, nil // Mengembalikan true jika count > 0, else false
}

// CreateComment menyimpan comment baru ke dalam database
func (repo *commentRepository) CreateComment(comment *models.Comment) (models.Comment, error) {
	// Menyimpan comment ke dalam database
	if err := repo.db.Create(comment).Error; err != nil {
		return models.Comment{}, err
	}
	return *comment, nil
}
