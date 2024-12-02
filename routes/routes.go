package routes

import (
	"go-restful-blog-api/v2/auth/middleware"
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

	postRepo := repositories.NewPostRepository(config.DB)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	commentRepo := repositories.NewCommentRepository(config.DB)           // Menggunakan CommentRepository
	commentService := services.NewCommentService(commentRepo)             // Menggunakan CommentService
	commentController := controllers.NewCommentController(commentService) // Menggunakan CommentController

	// Menentukan route untuk POST /register yang mengarah ke RegisterUser pada userController
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	// Group route yang memerlukan autentikasi
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/posts", postController.GetPosts)
		protected.POST("/posts", postController.CreatePost)
		protected.GET("/posts/:id", postController.GetPostByID)
		protected.PUT("/posts/", postController.UpdatePost)
		protected.DELETE("/posts/:id", postController.DeletePost)

		protected.POST("/comments", commentController.CreateComment)
	}

	// Kembalikan router yang sudah dikonfigurasi
	return router
}
