package services

import (
	"errors"
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

func FindRecipeService() ([]models.Recipe, error) {
	recipes, err := repositories.FindRecipe()
	return recipes, err
}

func FindByIdRecipeService(id int) (models.Recipe, error) {
	recipe, err := repositories.FindByIdRecipe(id)
	return recipe, err
}

func CreateRecipeService(recipe models.Recipe) {
	repositories.CreateRecipe(recipe)
}

func SaveRepiceService(id int, recipe models.Recipe) (models.Recipe, error) {
	response, err := repositories.SaveRepice(id, recipe)
	return response, err
}

func DeleteRecipeService(id int) (models.Recipe, error) {
	recipe, err := FindByIdRecipeService(id)

	if recipe.ID == 0 {
		return recipe, errors.New("Recipe not found")
	}

	repositories.DeleteRecipe(id)
	return recipe, err
}
