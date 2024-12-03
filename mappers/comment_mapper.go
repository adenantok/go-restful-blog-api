package mappers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/models"
)

// MapToComment mengonversi data dari CommentDTO ke Comment model
func MapToComment(commentDTO dto.CommentDTO) models.Comment {
	return models.Comment{
		ID:      commentDTO.ID,
		PostID:  commentDTO.PostID,
		UserID:  commentDTO.UserID,
		Content: commentDTO.Content,
	}
}

// MapToCommentDTO mengonversi data dari Comment model ke CommentDTO
func MapToCommentDTO(comment models.Comment) dto.CommentDTO {
	return dto.CommentDTO{
		ID:      comment.ID,
		PostID:  comment.PostID,
		UserID:  comment.UserID,
		Content: comment.Content,
	}
}

func MapToCreateComment(CreateCommentDTO dto.CreateCommentDTO) models.Comment {
	return models.Comment{
		PostID:  CreateCommentDTO.PostID,
		UserID:  CreateCommentDTO.UserID,
		Content: CreateCommentDTO.Content,
	}
}

func MapToCommentDTOs(comments []models.Comment) []dto.CommentDTO {
	commentDTOs := make([]dto.CommentDTO, len(comments))
	for i, comment := range comments {
		commentDTOs[i] = MapToCommentDTO(comment)
	}
	return commentDTOs
}
