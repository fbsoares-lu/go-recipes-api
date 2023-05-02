package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/middlewares"

	"github.com/gin-gonic/gin"
)

func IngredientRoutes(r *gin.Engine) {
	routes := r.Group("/api/ingredients")

	routes.POST("/", middlewares.RequiredAuth, controllers.CreateIngredient)
	routes.GET("/", middlewares.RequiredAuth, controllers.FindIngredient)
}
