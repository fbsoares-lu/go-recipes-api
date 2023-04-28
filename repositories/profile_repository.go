package repositories

import (
	"fbsoares-lu/go-recipes-api/models"
	"time"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindById(id int) (*models.Profile, error)
	Create(profile models.Profile) error
	Update(id int, profile models.Profile) error
	Delete(id int) error
}

type GORMProfileRepository struct {
	DB *gorm.DB
}

func (r *GORMProfileRepository) FindById(id int) (*models.Profile, error) {
	var profile models.Profile
	err := r.DB.First(&profile, id).Where("deleted_at", nil).Preload("File").Find(&profile).Error
	return &profile, err
}

func (r *GORMProfileRepository) Create(profile models.Profile) error {
	err := r.DB.Create(&profile).Error
	return err
}

func (r *GORMProfileRepository) Update(id int, profile models.Profile) error {
	err := r.DB.Model(&profile).UpdateColumns(profile).Error
	return err
}

func (r *GORMProfileRepository) Delete(id int) error {
	var profile models.Profile
	if err := r.DB.First(&profile).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&profile).UpdateColumn("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
