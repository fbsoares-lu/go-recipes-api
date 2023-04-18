package database

import "fbsoares-lu/go-recipes-api/models"

func SyncMigrations() {
	DB.AutoMigrate(&models.User{})
}
