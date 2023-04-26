package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/middlewares"
	"fbsoares-lu/go-recipes-api/repositories"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	routes := r.Group("/api/users")

	userRepository := &repositories.GORMUserRepository{DB: database.DB}
	userService := &services.UserService{UserRepository: userRepository}
	userController := &controllers.UserController{UserService: *userService}

	routes.POST("/signup", userController.Signup)
	routes.POST("/login", userController.Login)
	routes.GET("/", middlewares.RequiredAuth, func(c *gin.Context) {
		user, _ := c.Get("user")
		// user.(models.User).Name
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	})
}
