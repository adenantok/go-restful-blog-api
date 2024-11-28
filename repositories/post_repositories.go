package repositories

import (
	"go-restful-blog-api/v2/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPostByID(ID int) (models.Post, error)
	UpdatePost(post *models.Post) (models.Post, error)
	DeletePost(ID int) error
	GetPostsWithPagination(page, limit int) ([]models.Post, error)
	GetTotalRecords() (int, error)
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

func (repo *postRepository) UpdatePost(post *models.Post) (models.Post, error) {

	if err := repo.db.Save(&post).Error; err != nil {
		return models.Post{}, err
	}
	return *post, nil

}

func (repo *postRepository) DeletePost(ID int) error {

	if err := repo.db.Delete(&models.Post{}, ID).Error; err != nil {
		return err
	}
	return nil
}

func (repo *postRepository) GetPostsWithPagination(page, limit int) ([]models.Post, error) {
	var posts []models.Post
	offset := (page - 1) * limit
	if err := repo.db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *postRepository) GetTotalRecords() (int, error) {
	var count int64
	if err := repo.db.Model(&models.Post{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
