package mappers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/models"
)

func MapToPost(postDTO dto.PostDTO) models.Post {
	return models.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content, // Password akan dihash sebelum disimpan di database
	}
}

func MapToPostDTO(post models.Post) *dto.PostDTO {
	return &dto.PostDTO{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}
