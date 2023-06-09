package routes

import (
	"fbsoares-lu/go-recipes-api/controllers"
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/middlewares"
	"fbsoares-lu/go-recipes-api/repositories"
	"fbsoares-lu/go-recipes-api/services"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine) {
	routes := r.Group("/api/profiles")

	profileRepository := &repositories.GORMProfileRepository{DB: database.DB}
	profileService := &services.ProfileService{ProfileRepository: profileRepository}
	profileController := &controllers.ProfileController{ProfileService: *profileService}

	routes.POST("/", middlewares.RequiredAuth, profileController.CreateProfile)
	routes.GET("/:id", middlewares.RequiredAuth, profileController.FindOneProfile)
	routes.PUT("/:id", middlewares.RequiredAuth, profileController.UpdateProfile)
	routes.DELETE("/:id", middlewares.RequiredAuth, profileController.DeleteProfile)
}
