package repositories

import (
	"fbsoares-lu/go-recipes-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(name string, email string, password string, role string) error
}

type GORMUserRepository struct {
	DB *gorm.DB
}

func (r *GORMUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GORMUserRepository) Create(name string, email string, password string, role string) error {
	user := models.User{Name: name, Email: email, Password: password, Role: role}

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
