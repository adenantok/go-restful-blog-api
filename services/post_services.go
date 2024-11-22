package services

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/mappers"
	"go-restful-blog-api/v2/repositories"
)

type PostService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(postDTO *dto.PostDTO) (dto.PostDTO, error) {
	post := mappers.MapToPost(*postDTO)
	// Menyimpan post ke database melalui repository
	createdPost, err := s.repo.CreatePost(&post)
	if err != nil {
		return dto.PostDTO{}, err
	}

	// Mengonversi post yang baru dibuat ke PostDTO untuk response
	postDTO = mappers.MapToPostDTO(createdPost)

	return *postDTO, nil
}

func (s *PostService) GetPosts() ([]dto.PostDTO, error) {
	posts, err := s.repo.GetPosts()
	if err != nil {
		return nil, err
	}

	// postDTO:=mappers.MapToPostDTO(posts)
	// return postDTO,nil
	return mappers.MapToPostDTOs(posts), nil
}
