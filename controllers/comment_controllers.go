package controllers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/services"
	"go-restful-blog-api/v2/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service *services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{
		service: service,
	}
}

func (controller *CommentController) CreateComment(c *gin.Context) {
	idParam := c.Param("postID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID format")
		return
	}
	var createCommentDTO dto.CreateCommentDTO
	log.Println("postid", id)
	// Bind the incoming JSON data to CreateCommentDTO
	if err := c.ShouldBindJSON(&createCommentDTO); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		utils.InternalServerErrorResponse(c, "UserID not found in context")
		return
	}
	createCommentDTO.PostID = id
	createCommentDTO.UserID = userID.(int)
	log.Println("postid setelah", createCommentDTO.PostID)
	log.Println("userid", createCommentDTO.UserID)
	// Call the service to create a comment
	createdCommentDTO, err := controller.service.CreateComment(createCommentDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Return success response with created data
	utils.CreatedResponse(c, "Comment created successfully", createdCommentDTO)
}

func (controller *CommentController) GetCommentsByPostID(c *gin.Context) {
	// Mendapatkan parameter postID dari URL
	idParam := c.Param("postID")

	// Mengonversi postID dari string ke integer
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid post ID format")
		return
	}

	// Memanggil service untuk mendapatkan komentar berdasarkan postID
	comments, err := controller.service.GetCommentsByPostID(postID)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Mengembalikan respons sukses dengan data komentar
	utils.SuccessResponse(c, "Comments retrieved successfully", comments)
}
