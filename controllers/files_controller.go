package controllers

import (
	"fbsoares-lu/go-recipes-api/services"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	services.UploadFile(c)
}
