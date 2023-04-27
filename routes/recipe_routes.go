package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/middlewares"
	"fbsoares-lu/go-recipes-api/repositories"
	"fbsoares-lu/go-recipes-api/services"

	"github.com/gin-gonic/gin"
)

func RecipeRoutes(r *gin.Engine) {
	routes := r.Group("/api/recipes")

	recipeRepository := &repositories.GORMRecipeRepository{DB: database.DB}
	recipeService := &services.RecipeService{RecipeRepository: recipeRepository}
	recipeController := &controllers.RecipeController{RecipeService: *recipeService}

	routes.GET("/", middlewares.RequiredAuth, recipeController.FindRecipe)
	routes.GET("/:id", middlewares.RequiredAuth, recipeController.FindOneRecipe)
	routes.POST("/", middlewares.RequiredAuth, recipeController.CreateRecipe)
	routes.PUT("/:id", middlewares.RequiredAuth, recipeController.UpdateRecipe)
	routes.DELETE("/:id", middlewares.RequiredAuth, recipeController.DeleteRecipe)
}
