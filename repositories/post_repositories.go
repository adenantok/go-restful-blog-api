package repositories

import (
	"go-restful-blog-api/v2/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPostByID(ID int) (models.Post, error)
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

func (repo *postRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := repo.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *postRepository) GetPostByID(ID int) (models.Post, error) {
	var post models.Post
	if err := repo.db.First(&post, ID).Error; err != nil {
		return models.Post{}, err
	}
	return post, nil
}
