package dto

// CommentDTO digunakan untuk mengirim atau menerima data komentar dalam format JSON
type CommentDTO struct {
	ID      int    `json:"id" binding:"required"`
	PostID  int    `json:"post_id" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateCommentDTO struct {
	PostID  int    `json:"post_id" binding:"required"`
	UserID  int    `json:"user_id" `
	Content string `json:"content" binding:"required"`
}
