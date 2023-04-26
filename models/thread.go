package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Description string
	RecipeID    int
	Recipe      Recipe
	UserID      int
	User        User
}
