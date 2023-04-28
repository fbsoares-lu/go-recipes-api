package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")
	routes := r.Group("/api")

	routes.POST("/uploads", controllers.Upload)
}
