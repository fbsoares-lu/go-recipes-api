package services

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

func FindIngredientService() []models.Ingredient {
	ingredients := repositories.FindIngredient()
	return ingredients
}

func FindByIdIngredientService(id int) models.Ingredient {
	ingredient := repositories.FindByIdIngredient(id)
	return ingredient
}

func CreateIngredientService(ingredient models.Ingredient) {
	repositories.CreateIngredient(ingredient)
}

func SaveIngredientService(id int, ingredient models.Ingredient) models.Ingredient {
	repositories.SaveIngredient(id, ingredient)
	return ingredient
}

func DeleteIngredientService(id int) {
	repositories.DeleteIngredient(id)
}
