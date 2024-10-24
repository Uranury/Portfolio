package db

import (
	"REST-API/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.Employee{}, &models.Department{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database from GORM: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("failed to close database: %v", err)
	}
}
