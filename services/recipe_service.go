package services

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

func FindRecipeService() ([]models.Recipe, error) {
	recipes, err := repositories.FindRecipe()
	return recipes, err
}

func FindByIdRecipeService(id int) models.Recipe {
	recipe := repositories.FindByIdRecipe(id)
	return recipe
}

func CreateRecipeService(recipe models.Recipe) {
	repositories.CreateRecipe(recipe)
}

func SaveRepiceService(id int, recipe models.Recipe) models.Recipe {
	repositories.SaveRepice(id, recipe)
	return recipe
}

func DeleteRecipeService(id int) {
	repositories.DeleteRecipe(id)
}
