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

func (s *PostService) CreatePost(CreatePostDTO dto.CreatePostDTO) (dto.PostDTO, error) {
	post := mappers.MapToCreatePost(CreatePostDTO)
	// Menyimpan post ke database melalui repository
	createdPost, err := s.repo.CreatePost(&post)
	if err != nil {
		return dto.PostDTO{}, err
	}

	// Mengonversi post yang baru dibuat ke PostDTO untuk response
	postDTO := mappers.MapToPostDTO(createdPost)

	return postDTO, nil
}

func (s *PostService) GetPosts(page, limit int) ([]dto.PostDTO, int, error) {
	// Hitung total record untuk pagination
	totalRecords, err := s.repo.GetTotalRecords()
	if err != nil {
		return nil, 0, err
	}

	// Ambil data posts berdasarkan pagination
	posts, err := s.repo.GetPostsWithPagination(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Convert ke PostDTO
	postDTOs := mappers.MapToPostDTOs(posts)

	return postDTOs, totalRecords, nil
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

func (s *PostService) DeletePost(ID int, postDTO dto.PostDTO) error {

	existingData, err := s.repo.GetPostByID(ID)
	//log.Println("id", post.ID)
	if err != nil {
		return err // Kembalikan error jika data tidak ditemukan
	}

	// log.Println("id", existingData.UserID)
	// log.Println("id", postDTO.UserID)

	// Validasi bahwa user yang ingin menghapus adalah pemilik postingan
	if existingData.UserID != postDTO.UserID {
		return errors.New("unauthorized: user does not own this post")
	}

	if err := s.repo.DeletePost(ID); err != nil {
		return err
	}
	return nil
}
