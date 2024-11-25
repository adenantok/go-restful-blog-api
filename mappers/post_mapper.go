package mappers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/models"
)

func MapToPost(postDTO dto.PostDTO) models.Post {
	return models.Post{
		UserID:  postDTO.UserID,
		Title:   postDTO.Title,
		Content: postDTO.Content,
	}
}

func MapToPostDTO(post models.Post) dto.PostDTO {
	return dto.PostDTO{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func MapToPostDTOs(posts []models.Post) []dto.PostDTO {
	postDTOs := make([]dto.PostDTO, len(posts))
	for i, post := range posts {
		postDTOs[i] = dto.PostDTO{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
		}
	}
	return postDTOs
}
