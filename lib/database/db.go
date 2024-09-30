package database

import (
	"boilerplate/lib/database/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseWithDSN(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connection established")

	// AutoMigrate entities
	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
