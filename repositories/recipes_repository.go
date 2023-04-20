package repositories

import (
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/models"
	"time"
)

func FindRecipe() ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := database.DB.Find(&recipes).Where("deleted_at", nil).Preload("Ingredients").Find(&recipes).Error
	return recipes, err
}

func FindByIdRecipe(id int) models.Recipe {
	var recipe models.Recipe
	database.DB.First(&recipe, id)
	return recipe
}

func CreateRecipe(recipe models.Recipe) {
	database.DB.Create(&recipe)
}

func SaveRepice(id int, recipe models.Recipe) models.Recipe {
	database.DB.First(&recipe, id)
	database.DB.Model(&recipe).UpdateColumns(recipe)
	return recipe
}

func DeleteRecipe(id int) {
	var recipe models.Recipe
	database.DB.First(&recipe, id)
	database.DB.Model(&recipe).Update("deleted_at", time.Now())
}
