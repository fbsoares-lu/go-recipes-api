package database

import "fbsoares-lu/go-recipes-api/models"

func SyncMigrations() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.File{})
	DB.AutoMigrate(&models.Recipe{})
	DB.AutoMigrate(&models.Ingredient{})
}
