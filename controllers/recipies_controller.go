package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindRecipe(c *gin.Context) {
	recipes, err := services.FindRecipeService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func FindOneRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	recipe, err := services.FindByIdRecipeService(formatId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	services.CreateRecipeService(recipe)
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	var recipe models.Recipe

	updatedRecipe, err := services.SaveRepiceService(formatId, recipe)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, updatedRecipe)
}

func DeleteRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	formatId, _ := strconv.Atoi(recipeId)

	_, err := services.DeleteRecipeService(formatId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
