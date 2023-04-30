package middlewares

import (
	"fbsoares-lu/go-recipes-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequiredUserProfileMatched(c *gin.Context) {
	profileId, _ := strconv.Atoi(c.Params.ByName("id"))
	user, _ := c.Get("user")

	payloadProfile := user.(models.User).Profile.ID
	userHasProfile := payloadProfile == 0

	if userHasProfile {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if profileId != int(payloadProfile) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
