package repositories

import (
	"fbsoares-lu/go-recipes-api/models"
	"time"

	"gorm.io/gorm"
)

type RecipeRepository interface {
	Find() (*[]models.Recipe, error)
	FindById(id int) (*models.Recipe, error)
	Create(recipe models.Recipe) error
	Save(id int, recipe models.Recipe) (*models.Recipe, error)
	Delete(id int) error
}

type GORMRecipeRepository struct {
	DB *gorm.DB
}

func (r *GORMRecipeRepository) Find() (*[]models.Recipe, error) {
	var recipes []models.Recipe
	err := r.DB.Find(&recipes).Where("deleted_at", nil).Preload("User").Find(&recipes).Error
	return &recipes, err
}

func (r *GORMRecipeRepository) FindById(id int) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.DB.First(&recipe, id).Where("deleted_at", nil).Preload("Ingredients").Preload("Files").Find(&recipe).Error
	return &recipe, err
}

func (r *GORMRecipeRepository) Create(recipe models.Recipe) error {
	err := r.DB.Create(&recipe).Error
	return err
}

func (r *GORMRecipeRepository) Save(id int, recipe models.Recipe) (*models.Recipe, error) {
	err := r.DB.Model(&recipe).UpdateColumns(recipe).Error
	return &recipe, err
}

func (r *GORMRecipeRepository) Delete(id int) error {
	var recipe models.Recipe

	if err := r.DB.First(&recipe, id).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&recipe).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
