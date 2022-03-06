package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "postgres://postgres:mehran8282@localhost:5432/book_db"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrateBookErr := database.AutoMigrate(&Book{})
	if migrateBookErr != nil {
		panic(migrateBookErr)
	}
	migratePDFErr := database.AutoMigrate(&PDF{})
	if migratePDFErr != nil {
		panic(migratePDFErr)
	}

	DB = database
}
