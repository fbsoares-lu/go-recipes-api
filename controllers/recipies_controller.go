package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"

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

func FindOneRecipe() {}

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

func UpdateRecipe() {}

func DeleteRecipe() {}
