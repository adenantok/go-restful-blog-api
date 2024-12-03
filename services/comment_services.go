package services

import (
	"errors"
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/mappers"
	"go-restful-blog-api/v2/repositories"
	"log"
)

type CommentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) CreateComment(CreateCommentDTO dto.CreateCommentDTO) (dto.CommentDTO, error) {
	// Menggunakan PostID dari CreateCommentDTO langsung untuk validasi
	postID := CreateCommentDTO.PostID

	// Validasi apakah ID postingan valid
	postExists, err := s.repo.CheckPostExists(postID)
	if err != nil {
		return dto.CommentDTO{}, err
	}
	if !postExists {
		return dto.CommentDTO{}, errors.New("post ID does not exist")
	}

	// Map CreateCommentDTO ke model Comment
	comment := mappers.MapToCreateComment(CreateCommentDTO)
	log.Printf("userid %d", comment.UserID)
	// Simpan komentar ke database melalui repository
	createdComment, err := s.repo.CreateComment(&comment)
	if err != nil {
		return dto.CommentDTO{}, err
	}

	// Map model Comment ke CommentDTO untuk output
	commentDTO := mappers.MapToCommentDTO(createdComment)

	return commentDTO, nil
}

func (s *CommentService) GetCommentsByPostID(postID int) ([]dto.CommentDTO, error) {
	// Mengambil data komentar berdasarkan postID dari repository
	comments, err := s.repo.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}

	// Mengonversi hasil model ke DTO menggunakan mapper
	commentDTOs := mappers.MapToCommentDTOs(comments)

	return commentDTOs, nil
}
