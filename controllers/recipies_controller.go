package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	RecipeService services.RecipeService
}

func (controller *RecipeController) FindRecipe(c *gin.Context) {
	recipes, err := controller.RecipeService.FindRecipeService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func (controller *RecipeController) FindOneRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	recipe, err := controller.RecipeService.FindByIdRecipeService(formatId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (controller *RecipeController) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	user, _ := c.Get("user")

	userId := user.(models.User).ID

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := controller.RecipeService.CreateRecipeService(recipe, int(userId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *RecipeController) UpdateRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	var recipe models.Recipe

	updatedRecipe, err := controller.RecipeService.SaveRepiceService(formatId, recipe)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, updatedRecipe)
}

func (controller *RecipeController) DeleteRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	err := controller.RecipeService.DeleteRecipeService(formatId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
