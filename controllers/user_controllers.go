package controllers

import (
	"go-restful-blog-api/v2/dto"
	"go-restful-blog-api/v2/services"
	"go-restful-blog-api/v2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// userController adalah struct yang menangani request terkait pengguna
type userController struct {
	service *services.UserService
}

// NewUserController membuat instance baru dari userController
func NewUserController(service *services.UserService) *userController {
	return &userController{
		service: service,
	}
}

// RegisterUser adalah handler untuk mendaftar user baru
func (controller *userController) RegisterUser(c *gin.Context) {
	var userDTO dto.UserDTO

	// Bind data dari request body ke UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		// Jika data yang diterima tidak valid (misalnya field yang wajib tidak ada)
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Panggil service untuk register user
	user, err := controller.service.RegisterUser(userDTO)
	if err != nil {
		// Jika ada error dari service (misalnya username sudah ada)
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Kembalikan response sukses dengan data user
	utils.CreatedResponse(c, "User registered successfully", user)
	//c.JSON(http.StatusOK, gin.H{"user": user})
}

// LoginUser adalah handler untuk login user
func (controller *userController) LoginUser(c *gin.Context) {
	var loginDTO dto.UserDTO

	// Bind data dari request body ke LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		// Jika data yang diterima tidak valid (misalnya field yang wajib tidak ada)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk melakukan login
	user, token, err := controller.service.LoginUser(loginDTO)
	if err != nil {
		// Jika terjadi error pada service (misalnya username atau password salah)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response sukses dengan data user dan token
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
