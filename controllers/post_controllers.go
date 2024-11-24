package controllers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/services"
	"go-restful-blog-api/v2/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postController struct {
	service *services.PostService
}

func NewPostController(service *services.PostService) *postController {
	return &postController{
		service: service,
	}
}

// HandleCreatePost handles the creation of a new post
func (controller *postController) CreatePost(c *gin.Context) {
	var postDTO dto.PostDTO

	// Bind the incoming JSON data to postDTO
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		utils.InternalServerErrorResponse(c, "UserID not found in context")
		return
	}

	postDTO.UserID = userID.(int)

	// Pass the DTO to the PostService to process the logic
	createdPost, err := controller.service.CreatePost(postDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Return success response with created data
	utils.CreatedResponse(c, "Post created successfully", createdPost)
}

func (controller *postController) GetPosts(c *gin.Context) {
	posts, err := controller.service.GetPosts()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "posts retrieved successfully", posts)
}

func (controller *postController) GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID format")
		return
	}
	post, err := controller.service.GetPostByID(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "post retrieved successfully", post)
}

func (controller *postController) UpdatePost(c *gin.Context) {
	var postDTO dto.PostDTO

	// Bind the incoming JSON data to postDTO
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		utils.InternalServerErrorResponse(c, "UserID not found in context")
		return
	}

	postDTO.UserID = userID.(int)

	// Pass the DTO to the PostService to process the logic
	updatePost, err := controller.service.UpdatePost(postDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Return success response with created data
	utils.CreatedResponse(c, "Post updated successfully", updatePost)
}

func (controller *postController) DeletePost(c *gin.Context) {
	var postDTO dto.PostDTO
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID format")
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		utils.InternalServerErrorResponse(c, "UserID not found in context")
		return
	}

	postDTO.UserID = userID.(int)

	err = controller.service.DeletePost(id, postDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "post deleted successfully", nil)
}
