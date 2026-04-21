package db

import (
	"log"
	"pool/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// The Postgres connection string 
	dsn := "host=localhost user=postgres password=postgres dbname=test-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Connect using the Postgres driver
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		// Use log.Fatalf instead of panic for cleaner exit messages
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Successfully connected to Postgres!")

	// AutoMigrate models
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
