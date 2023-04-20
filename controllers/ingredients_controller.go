package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindIngredient(c *gin.Context) {
	ingredients := services.FindIngredientService()
	c.JSON(http.StatusOK, ingredients)
}

func FindOneIngredient() {}

func CreateIngredient(c *gin.Context) {
	var ingredient models.Ingredient

	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	services.CreateIngredientService(ingredient)
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateIngredient() {}

func DeleteIngredient() {}
