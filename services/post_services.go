package services

import (
	"errors"
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

func (s *PostService) CreatePost(postDTO dto.PostDTO) (dto.PostDTO, error) {
	post := mappers.MapToPost(postDTO)
	// Menyimpan post ke database melalui repository
	createdPost, err := s.repo.CreatePost(&post)
	if err != nil {
		return dto.PostDTO{}, err
	}

	// Mengonversi post yang baru dibuat ke PostDTO untuk response
	postDTO = mappers.MapToPostDTO(createdPost)

	return postDTO, nil
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

func (s *PostService) GetPostByID(ID int) (dto.PostDTO, error) {
	post, err := s.repo.GetPostByID(ID)
	if err != nil {
		return dto.PostDTO{}, err
	}
	postDTO := mappers.MapToPostDTO(post)
	return postDTO, nil
}

func (s *PostService) UpdatePost(postDTO dto.PostDTO) (dto.PostDTO, error) {
	post := mappers.MapToPost(postDTO)

	// Mendapatkan data existing berdasarkan ID
	existingData, err := s.repo.GetPostByID(post.ID)
	if err != nil {
		return dto.PostDTO{}, err // Kembalikan error jika data tidak ditemukan
	}

	// Validasi bahwa user yang ingin mengedit adalah pemilik postingan
	if existingData.UserID != postDTO.UserID {
		return dto.PostDTO{}, errors.New("unauthorized: user does not own this post")
	}

	if post.Content != "" {
		existingData.Content = post.Content
	}
	if post.Title != "" {
		existingData.Title = post.Title
	}

	updatedPost, err := s.repo.UpdatePost(&existingData)
	if err != nil {
		return dto.PostDTO{}, nil
	}
	postDTO = mappers.MapToPostDTO(updatedPost)
	return postDTO, nil
}
