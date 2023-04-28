package routes

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	UserRoutes(r)
	RecipeRoutes(r)
	ProfileRoutes(r)
	IngredientRoutes(r)
	FileRoutes(r)

	r.Run(":5000")
}
