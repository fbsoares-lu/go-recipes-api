package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	RecipeID    uint
}
