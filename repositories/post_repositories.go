package repositories

import (
	"go-restful-blog-api/v2/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) (models.Post, error)
	//GetUserByUsername(username string) (models.User, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (repo *postRepository) CreatePost(post *models.Post) (models.Post, error) {
	// Menyimpan post ke dalam database
	if err := repo.db.Create(post).Error; err != nil {
		return models.Post{}, err
	}
	return *post, nil
}
