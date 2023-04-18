package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/users", middlewares.RequiredAuth, func(c *gin.Context) {
		user, _ := c.Get("user")

		// OBS
		// user.(models.User).Name

		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	})

	r.Run(":5000")
}
