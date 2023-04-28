package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService services.ProfileService
}

func (controller *ProfileController) FindOneProfile(c *gin.Context) {
	profileId, _ := strconv.Atoi(c.Params.ByName("id"))

	recipe, err := controller.ProfileService.FindOneProfileService(profileId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (controller *ProfileController) CreateProfile(c *gin.Context) {
	var profile models.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := controller.ProfileService.CreateProfileService(profile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *ProfileController) UpdateProfile(c *gin.Context) {
	profileId, _ := strconv.Atoi(c.Params.ByName("id"))

	var profile models.Profile

	updatedProfile, err := controller.ProfileService.SaveProfileService(profileId, profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, updatedProfile)
}

func (controller *ProfileController) DeleteProfile(c *gin.Context) {
	profileId, _ := strconv.Atoi(c.Params.ByName("id"))

	err := controller.ProfileService.DeleteProfileService(profileId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
