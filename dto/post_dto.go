package dto

// CreatePostDTO adalah DTO untuk membuat post baru
type CreatePostDTO struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title" binding:"required"`   // Validasi agar title tidak kosong
	Content string `json:"content" binding:"required"` // Validasi agar content tidak kosong
}

type PostDTO struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostDTOResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
