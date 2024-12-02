package controllers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/services"
	"go-restful-blog-api/v2/utils"
	"log"

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
	var createCommentDTO dto.CreateCommentDTO
	log.Println("userid", createCommentDTO.UserID)
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

	createCommentDTO.UserID = userID.(int)
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
