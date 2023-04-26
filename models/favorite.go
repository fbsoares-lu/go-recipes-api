package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID   int
	User     User
	RecipeID int
	Recipe   Recipe
}
