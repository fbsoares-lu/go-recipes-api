package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	// Files
	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")
	r.POST("/api/uploads", controllers.Upload)

	// Authentication
	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)

	// Users
	r.GET("/api/users", middlewares.RequiredAuth, func(c *gin.Context) {
		user, _ := c.Get("user")
		// user.(models.User).Name
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	})

	r.Run(":5000")
}
