package repositories

import (
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/models"
	"time"
)

func FindIngredient() []models.Ingredient {
	var ingredients []models.Ingredient
	database.DB.Find(&ingredients).Where("deleted_at", nil)
	return ingredients
}

func FindByIdIngredient(id int) models.Ingredient {
	var ingredient models.Ingredient
	database.DB.First(&ingredient, id)
	return ingredient
}

func CreateIngredient(ingredient models.Ingredient) {
	database.DB.Create(&ingredient)
}

func SaveIngredient(id int, ingredient models.Ingredient) models.Ingredient {
	database.DB.First(&ingredient, id)
	database.DB.Model(&ingredient).UpdateColumns(ingredient)
	return ingredient
}

func DeleteIngredient(id int) {
	var ingredient models.Ingredient
	database.DB.First(&ingredient, id)
	database.DB.Model(&ingredient).Update("deleted_at", time.Now())
}
