package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/middlewares"

	"github.com/gin-gonic/gin"
)

func IngredientRoutes(r *gin.Engine) {
	routes := r.Group("/api/ingredients")

	routes.POST("/api/ingredients", middlewares.RequiredAuth, controllers.CreateIngredient)
	routes.GET("/api/ingredients", middlewares.RequiredAuth, controllers.FindIngredient)
}
