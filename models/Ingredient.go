package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Description string
	Recipes     []*Recipe `gorm:"many2many:recipe_ingredients;"`
}
