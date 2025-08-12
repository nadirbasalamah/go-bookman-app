package storage

import (
	"fmt"
	"go-bookman-app/models"
	"go-bookman-app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_USER"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_NAME"),
		utils.GetConfig("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return nil, err
	}

	return db, nil
}
