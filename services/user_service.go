package services

import (
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func FindByEmailUserService(email string, c *gin.Context) models.User {
	user := repositories.FindByEmail(email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	return user
}

func CreateUserService(user models.User, c *gin.Context) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	result := repositories.Create(user.Name, user.Email, string(hash), user.Role)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Falied to create user",
		})
	}
}

func DecodePassword(hasedPassword string, password string, c *gin.Context) {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed comparation",
		})
	}
}

func Authentication(email string, password string, c *gin.Context) {
	user := FindByEmailUserService(email, c)
	DecodePassword(user.Password, password, c)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed comparation",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
}
