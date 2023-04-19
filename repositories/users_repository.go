package repositories

import (
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/models"

	"gorm.io/gorm"
)

func Create(name string, email string, password string, role string) *gorm.DB {
	user := models.User{Name: name, Email: email, Password: password, Role: role}
	result := database.DB.Create(&user)

	return result
}

func FindByEmail(email string) models.User {
	var user models.User
	database.DB.First(&user, "email = ?", email)
	return user
}
