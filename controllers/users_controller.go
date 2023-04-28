package controllers

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (uc *UserController) Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// if err := models.User{}.Validate(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	err := uc.UserService.CreateUserService(user)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (uc *UserController) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	token, err := uc.UserService.Authentication(body.Email, body.Password)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
