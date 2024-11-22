package main

import (
	"go-restful-blog-api/v2/config"
	"go-restful-blog-api/v2/models"
	"go-restful-blog-api/v2/routes"
	"log"
)

func main() {
	config.ConnectDB()
	//config.TestDBConnection()
	models.MigrateUser(config.DB)
	models.MigratePost(config.DB)

	// Menyiapkan router
	router := routes.SetupRouter()

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server gagal dijalankan: ", err)
	}
}
