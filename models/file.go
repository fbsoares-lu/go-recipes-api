package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	OriginalName string
	RecipeID     uint
	ProfileID    uint `gorm:"uniqueIndex"`
}
