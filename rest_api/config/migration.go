package config

import "rest/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Transaction{})
}
