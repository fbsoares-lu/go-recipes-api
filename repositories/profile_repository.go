package repositories

import (
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/models"
	"time"
)

func FindByIdProfile(id int) (models.Profile, error) {
	var profile models.Profile
	err := database.DB.First(&profile, id).Error
	return profile, err
}

func CreateProfile(profile models.Profile) {
	database.DB.Create(&profile)
}

func UpdateProfile(id int, profile models.Profile) (models.Profile, error) {
	database.DB.First(&profile, id)
	err := database.DB.Model(&profile).UpdateColumns(profile).Error
	return profile, err
}

func DeleteProfile(id int) {
	var profile models.Profile
	database.DB.First(&profile)
	database.DB.Model(&profile).UpdateColumn("deleted_at", time.Now())
}
