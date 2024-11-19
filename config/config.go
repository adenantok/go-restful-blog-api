package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user=postgres password=indonesia dbname=go-restful-blog-api host=localhost port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		os.Exit(1)
	}
}

// TestDBConnection untuk menguji apakah koneksi ke database berhasil
func TestDBConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
		os.Exit(1)
	}

	// Cek koneksi dengan ping ke database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		os.Exit(1)
	}

	log.Println("Database connection is successful!")
}
