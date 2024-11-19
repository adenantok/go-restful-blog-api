package routes

import (
	"go-restful-blog-api/v2/config"
	"go-restful-blog-api/v2/controllers"  // Mengimpor controller yang sesuai
	"go-restful-blog-api/v2/repositories" // Mengimpor repository yang sesuai
	"go-restful-blog-api/v2/services"     // Mengimpor service yang sesuai

	"github.com/gin-gonic/gin" // Mengimpor Gin framework
)

func SetupRouter() *gin.Engine {
	//config.ConnectDB()

	// Membuat instance dari router Gin
	router := gin.Default()

	// Membuat instance dari repository, service, dan controller
	userRepo := repositories.NewUserRepository(config.DB)        // Membuat instance UserRepository
	userService := services.NewUserService(userRepo)             // Membuat instance UserService
	userController := controllers.NewUserController(userService) // Membuat instance UserController

	// Menentukan route untuk POST /register yang mengarah ke RegisterUser pada userController
	router.POST("/register", userController.RegisterUser)

	// Kembalikan router yang sudah dikonfigurasi
	return router
}