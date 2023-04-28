package services

import (
	"errors"
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

type RecipeService struct {
	RecipeRepository repositories.RecipeRepository
}

func (service *RecipeService) FindRecipeService() (*[]models.Recipe, error) {
	recipes, err := service.RecipeRepository.Find()

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (service *RecipeService) FindByIdRecipeService(id int) (*models.Recipe, error) {
	recipe, err := service.RecipeRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if recipe.ID == 0 {
		return nil, errors.New("recipe not found")
	}

	return recipe, nil
}

func (service *RecipeService) CreateRecipeService(recipe models.Recipe) error {
	err := service.RecipeRepository.Create(recipe)

	if err != nil {
		return err
	}

	return nil
}

func (service *RecipeService) SaveRepiceService(id int, recipe models.Recipe) (*models.Recipe, error) {
	currentRecipe, err := service.RecipeRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if currentRecipe.ID == 0 {
		return nil, errors.New("recipe not found")
	}

	response, err := service.RecipeRepository.Save(id, recipe)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *RecipeService) DeleteRecipeService(id int) error {
	recipe, err := service.RecipeRepository.FindById(id)

	if err != nil {
		return errors.New("failed to find recipe")
	}

	if recipe.ID == 0 {
		return errors.New("recipe not found")
	}

	err = service.RecipeRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
