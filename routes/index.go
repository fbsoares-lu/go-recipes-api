package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/middlewares"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")

	UserRoutes(r)
	RecipeRoutes(r)

	r.POST("/api/uploads", controllers.Upload)

	r.POST("/api/ingredients", middlewares.RequiredAuth, controllers.CreateIngredient)
	r.GET("/api/ingredients", middlewares.RequiredAuth, controllers.FindIngredient)

	r.Run(":5000")
}
