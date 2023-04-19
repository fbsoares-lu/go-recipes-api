package services

import (
	"context"
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/models"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload image",
		})
		return
	}

	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload image",
		})
		return
	}

	f, openErr := file.Open()

	if openErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("go-recipies"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload image to s3",
		})
		return
	}

	database.DB.Create(&models.File{
		OriginalName: result.Location,
	})
}
